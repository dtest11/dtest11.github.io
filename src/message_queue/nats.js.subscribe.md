nats js subscribe 模式实现

1. 客户端提供了如下接口

```go
func (js *js) subscribe(subj, queue string, cb MsgHandler, ch chan *Msg, isSync, isPullMode bool, opts []SubOpt) (*Subscription, error) {

```

2. nats-server

   1. parse.go

   2. processsub

       ```go
       processSubEx(subject, queue, bsid []byte, cb msgHandler, noForward, si, rsi bool) (*subscription
       ```

      

   3. 更新sublist 上到订阅对象

      ```go
      	es := c.subs[sid]
      	if es == nil {
      		c.subs[sid] = sub
      		if acc != nil && acc.sl != nil {
      			err = acc.sl.Insert(sub)// 重点是这个地方
      			if err != nil {
      				delete(c.subs, sid)
      			} else {
      				updateGWs = c.srv.gateway.enabled
      			}
      		}
      	}
      ```
下面我们来看消息是如何被客户端消费的


1.nats-server 的recover
nats server 在启动到时候会使用recover 来恢复所有到jetstream ，同时针对每一个subject 来注册一个handler,处理producer产生到消息


EnableJetStream =>

<= addstream

```go

	// Now recover the streams.
	fis, _ := ioutil.ReadDir(sdir)
	for _, fi := range fis {
	
		mset, err := a.addStream(&cfg.StreamConfig)
		if err != nil {
			s.Warnf("  Error recreating stream %q: %v", cfg.Name, err)
			continue
		}
		if !cfg.Created.IsZero() {
			mset.setCreatedTime(cfg.Created)
		}

		state := mset.state()
		s.Noticef("  Restored %s messages for stream %q", comma(int64(state.Msgs)), fi.Name())

	
	}
```

add consumers =>

```go
for _, e := range consumers {
		for _, ofi := range ofis {
			metafile := path.Join(e.odir, ofi.Name(), JetStreamMetaFile)
			obs, err := e.mset.addConsumer(&cfg.ConsumerConfig)
			
		}
	}
```

<= addConsumerWithAssignment
加载conumer的一些基本配置
```go
/ Set name, which will be durable name if set, otherwise we create one at random.
	o := &consumer{
		mset:    mset,
		js:      s.getJetStream(),
		acc:     a,
		srv:     s,
		client:  s.createInternalJetStreamClient(),
		sysc:    s.createInternalJetStreamClient(),
		cfg:     *config,
		dsubj:   config.DeliverSubject,
		outq:    mset.outq,
		active:  true,
		qch:     make(chan struct{}),
		mch:     make(chan struct{}, 1),
		sfreq:   int32(sampleFreq),
		maxdc:   uint64(config.MaxDeliver),
		maxp:    config.MaxAckPending,
		created: time.Now().UTC(),
	}
```
开启client 上线异步通知
```go
if config.Direct || (!s.JetStreamIsClustered() && s.standAloneMode()) {
		o.setLeader(true)
	}
```
<= setLeader

```go

o.inch = make(chan bool, 8)
logrus.Infof("consumer 注册通知 subject:%s",o.cfg.FilterSubject)
o.acc.sl.registerNotification(o.cfg.DeliverSubject, o.cfg.DeliverGroup, o.inch)
if o.active = <-o.inch; !o.active {
	// Check gateways in case they are enabled.
	if s.gateway.enabled {
		o.active = s.hasGatewayInterest(o.acc.Name, o.cfg.DeliverSubject)
		stopAndClearTimer(&o.gwdtmr)
		o.gwdtmr = time.AfterFunc(time.Second, func() { o.watchGWinterest() })
	}
} else {
	o.checkQueueInterest()
}
			

```
<= loopAndGatherMsgs
启动消息中心

```go

func (o *consumer) loopAndGatherMsgs(qch chan struct{}) {
	// On startup check to see if we are in a a reply situation where replay policy is not instant.
	var (
		lts  int64 // last time stamp seen, used for replay.
		lseq uint64
	)

	o.mu.Lock()
	s := o.srv
	// For idle heartbeat support.
	var hbc <-chan time.Time
	hbd, hb := o.hbTimer()
	if hb != nil {
		hbc = hb.C
	}
	// Interest changes.
	inch := o.inch
	o.mu.Unlock()

	// Deliver all the msgs we have now, once done or on a condition, we wait for new ones.
	for {
		// If we are in push mode and not active or under flowcontrol let's stop sending.
		if o.isPushMode() {
			if !o.active {
				goto waitForMsgs
			}
		}

		// If we are in pull mode and no one is waiting already break and wait.
		if o.isPullMode() && !o.checkWaitingForInterest() {
			goto waitForMsgs
		}

		subj, hdr, msg, seq, dc, ts, err = o.getNextMsg()

		// On error either wait or return.
		if err != nil {
			if err == ErrStoreMsgNotFound || err == ErrStoreEOF || err == errMaxAckPending {
				goto waitForMsgs
			} else {
				o.mu.Unlock()
				s.Errorf("Received an error looking up message for consumer: %v", err)
				return
			}
		}

		if wr := o.waiting.pop(); wr != nil {
			dsubj = wr.reply
		} else {
			dsubj = o.dsubj
		}

		// If we are in a replay scenario and have not caught up check if we need to delay here.
		if o.replay && lts > 0 {
			if delay = time.Duration(ts - lts); delay > time.Millisecond {
				o.mu.Unlock()
				select {
				case <-qch:
					return
				case <-time.After(delay):
				}
				o.mu.Lock()
			}
		}
		
		// Do actual delivery.
		o.deliverMsg(dsubj, subj, hdr, msg, seq, dc, ts)

		// Reset our idle heartbeat timer if set.
		if hb != nil {
			hb.Reset(hbd)
		}

		o.mu.Unlock()
		continue

	waitForMsgs:
		// If we were in a replay state check to see if we are caught up. If so clear.
		if o.replay && o.sseq > lseq {
			o.replay = false
		}

		// We will wait here for new messages to arrive.
		mch, outq, odsubj := o.mch, o.outq, o.cfg.DeliverSubject
		o.mu.Unlock()

		select {
		case interest := <-inch:
			// inch can be nil on pull-based, but then this will
			// just block and not fire.
			o.updateDeliveryInterest(interest)
		case <-qch:
			return
		case <-mch:
			// Messages are waiting.
		case <-hbc:
			if o.isActive() {
				const t = "NATS/1.0 100 Idle Heartbeat\r\n%s: %d\r\n%s: %d\r\n\r\n"
				sseq, dseq := o.lastDelivered()
				hdr := []byte(fmt.Sprintf(t, JSLastConsumerSeq, dseq, JSLastStreamSeq, sseq))
				if fcp := o.fcID(); fcp != _EMPTY_ {
					// Add in that we are stalled on flow control here.
					addOn := []byte(fmt.Sprintf("%s: %s\r\n\r\n", JSConsumerStalled, fcp))
					hdr = append(hdr[:len(hdr)-LEN_CR_LF], []byte(addOn)...)
				}
				outq.send(&jsPubMsg{odsubj, _EMPTY_, _EMPTY_, hdr, nil, nil, 0, nil})
			}
			// Reset our idle heartbeat timer.
			hb.Reset(hbd)
		}
	}
}
```




