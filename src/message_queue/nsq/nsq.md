

1.  NSQlookupd

记录整个服务topic. channel 对应的生产者的信息
```go


	router.Handle("GET", "/ping", http_api.Decorate(s.pingHandler, log, http_api.PlainText))
	router.Handle("GET", "/info", http_api.Decorate(s.doInfo, log, http_api.V1))

	// v1 negotiate
	router.Handle("GET", "/debug", http_api.Decorate(s.doDebug, log, http_api.V1))
	router.Handle("GET", "/lookup", http_api.Decorate(s.doLookup, log, http_api.V1))
	router.Handle("GET", "/topics", http_api.Decorate(s.doTopics, log, http_api.V1))
	router.Handle("GET", "/channels", http_api.Decorate(s.doChannels, log, http_api.V1))
	router.Handle("GET", "/nodes", http_api.Decorate(s.doNodes, log, http_api.V1))

	// only v1
	router.Handle("POST", "/topic/create", http_api.Decorate(s.doCreateTopic, log, http_api.V1))
	router.Handle("POST", "/topic/delete", http_api.Decorate(s.doDeleteTopic, log, http_api.V1))
	router.Handle("POST", "/channel/create", http_api.Decorate(s.doCreateChannel, log, http_api.V1))
	router.Handle("POST", "/channel/delete", http_api.Decorate(s.doDeleteChannel, log, http_api.V1))
	router.Handle("POST", "/topic/tombstone", http_api.Decorate(s.doTombstoneTopicProducer, log, http_api.V1))

	// debug
	router.HandlerFunc("GET", "/debug/pprof", pprof.Index)
	router.HandlerFunc("GET", "/debug/pprof/cmdline", pprof.Cmdline)
	router.HandlerFunc("GET", "/debug/pprof/symbol", pprof.Symbol)
	router.HandlerFunc("POST", "/debug/pprof/symbol", pprof.Symbol)
	router.HandlerFunc("GET", "/debug/pprof/profile", pprof.Profile)
	router.Handler("GET", "/debug/pprof/heap", pprof.Handler("heap"))
	router.Handler("GET", "/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handler("GET", "/debug/pprof/block", pprof.Handler("block"))
	router.Handler("GET", "/debug/pprof/threadcreate", pprof.Handler("threadcreate"))

```


2. nsqd 
维护 消费者 / 生产者的工作

3. 发送消息的过程
client => nsqd 走TCP connect
```go

// Connect dials and bootstraps the nsq connection
// (including IDENTIFY) and returns the IdentifyResponse
func (c *Conn) Connect() (*IdentifyResponse, error) {
	dialer := &net.Dialer{
		LocalAddr: c.config.LocalAddr,
		Timeout:   c.config.DialTimeout,
	}

	conn, err := dialer.Dial("tcp", c.addr)
	if err != nil {
		return nil, err
	}
	c.conn = conn.(*net.TCPConn)
	c.r = conn
	c.w = conn

	_, err = c.Write(MagicV2)
	if err != nil {
		c.Close()
		return nil, fmt.Errorf("[%s] failed to write magic - %s", c.addr, err)
	}

	resp, err := c.identify()
	if err != nil {
		return nil, err
	}

	if resp != nil && resp.AuthRequired {
		if c.config.AuthSecret == "" {
			c.log(LogLevelError, "Auth Required")
			return nil, errors.New("Auth Required")
		}
		err := c.auth(c.config.AuthSecret)
		if err != nil {
			c.log(LogLevelError, "Auth Failed %s", err)
			return nil, err
		}
	}

	c.wg.Add(2)
	atomic.StoreInt32(&c.readLoopRunning, 1)
	go c.readLoop()
	go c.writeLoop()
	return resp, nil
}

```
```go
Publish(topic string, body []byte) error 
MultiPublish(topic string, body [][]byte) error
DeferredPublish(topic string, delay time.Duration, body []byte)
DeferredPublishAsync(topic string, delay time.Duration, body []byte,
	doneChan chan *ProducerTransaction, args ...interface{}) error
 ```
 以上是nsqd 暴露出来发送消息的接口,他们都是通过
 ```go
sendCommandAsync
sendCommand(cmd *Command) error
 ```
这些方法来写入消息到 nsqd

sendCommand
||
sendCommand(cmd *Command) error

return result 

![这是图片](/message_queue/nsq/nsq.png) "Magic Gardens")

TOOD:
消息是否有序