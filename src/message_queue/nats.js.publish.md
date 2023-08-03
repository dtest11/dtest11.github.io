publish(nats.go)
```go
func (js *js) PublishMsg(m *Msg, opts ...PubOpt) (*PubAck, error) {
```

nats-server

. parse.go

. processpub


```go
func (c *client) processInboundClientMsg(msg []byte) (bool, bool) {

```

```go

didDeliver, qnames = c.processMsgResults(c.acc, r, msg, c.pa.deliver, c.pa.subject, c.pa.reply, flag)

```
推送到jetstream 系统

#### 由nats. subcribe 中我们知道stream 启动到时候，会根据topic 订阅一个msg handler 来启动一个内部到消费者来接受publish 的msg

```go
func (mset *stream) subscribeToStream() error {
	if mset.active {
		return nil
	}
	for _, subject := range mset.cfg.Subjects {
		logrus.Infof("subject:%s",subject)
		if _, err := mset.subscribeInternal(subject, mset.processInboundJetStreamMsg); err != nil {
			return err
		}
	}

```
## 流传 processJetStreamMsg
```go
func (mset *stream) processJetStreamMsg(subject, reply string, hdr, msg []byte, lseq uint64, ts int64) error {
```

检查是否有订阅者
```go
	if noInterest {
```
没有订阅者直接返回，然后存储消息
```go

// Store actual message_queue.
	if lseq == 0 && ts == 0 {
		seq, ts, err = store.StoreMsg(subject, hdr, msg)
	} else {
		// Make sure to take into account any message assignments that we had to skip (clfs).
		seq = lseq + 1 - clfs
		err = store.StoreRawMsg(subject, hdr, msg, seq, ts)
	}
```
如果有订阅者

```go
	if err == nil && seq > 0 && numConsumers > 0 {
		mset.mu.Lock()
		for _, o := range mset.consumers {
			o.mu.Lock()
			if o.isLeader() {
				if o.isFilteredMatch(subject) {
					o.sgap++
				}
				o.signalNewMessages()
				fmt.Println("signalNewMessages 有新消息进去")
			}
			o.mu.Unlock()
		}
		mset.mu.Unlock()
	}
```
向所有到订阅者发送一个channel 通知有最新到消息进去


