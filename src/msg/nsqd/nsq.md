

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


TOOD:
消息是否有序