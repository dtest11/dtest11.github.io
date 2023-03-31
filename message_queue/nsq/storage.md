# 存贮

![这是图片](/message_queue/nsq/storage.png)


1. 待确认的消息

	inFlightMessages map[MessageID]*Message
	inFlightPQ       inFlightPqueue //最大堆

2. 延迟消息【DPUB】
```go
	if chanMsg.deferred != 0 {
				channel.PutMessageDeferred(chanMsg, chanMsg.deferred)
				continue
			}

func (c *Channel) StartDeferredTimeout(msg *Message, timeout time.Duration) error {
	absTs := time.Now().Add(timeout).UnixNano()
	item := &pqueue.Item{Value: msg, Priority: absTs}
	err := c.pushDeferredMessage(item)
	if err != nil {
		return err
	}
	c.addToDeferredPQ(item)
	return nil
}

```
close 的时候刷盘