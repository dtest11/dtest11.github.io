# client(nats.go)
```go
		m, err = nc.newRequest(subj, hdr, data, timeout)

```
server:
```

readLoop => parse => processPub => output subject and reply 

func (c *client) processInboundMsg(msg []byte) {

```
