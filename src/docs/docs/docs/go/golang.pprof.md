### benchmark

### pprof
采集内存
 go tool pprof http://localhost:6060/debug/pprof/heap

采集CPU

go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

采集锁
runtime.SetBlockProfileRate 需要开启
go tool pprof http://localhost:6060/debug/pprof/block

runtime.SetMutexProfileFraction

go tool pprof http://localhost:6060/debug/pprof/mutex

采集trace
wget -O trace.out http://localhost:6060/debug/pprof/trace?seconds=5
go tool trace trace.out

https://go.dev/blog/pprof

```go
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		cache := make([]byte, 1*100) //100M
		cache[0] = 8
		fmt.Println("hello")
		writer.WriteHeader(200)
		_, _ = writer.Write(cache)
	})
	http.ListenAndServe(":6060", nil)
}

```

### trace

        
### network 
network [https://http2-explained.haxx.se/zh/part1]
