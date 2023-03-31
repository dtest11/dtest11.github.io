1. ### dst

   保证系统的可用性，不会被流量拖垮

2. ### status

   熔断器一般具有三个状态：

   1. 关闭：默认状态，请求能被到达目标服务，同时统计在窗口时间成功和失败次数，如果达到错误率阈值将会进入断开状态。

   2. 断开： 此状态下将会直接返回错误，如果有 fallback 配置则直接调用 fallback 方法。

   3. 半断开：进行断开状态会维护一个超市时间，到达超时时间开始进入 半断开 状态，尝试允许一部门请求正常通过并统计成功数量，如果请求正常则认为此时目标服务已恢复进入 关闭 状态，否则进入 断开 状态。半断开 状态存在的目的在于实现了自我修复，同时防止正在恢复的服务再次被大量打垮。

      作者：吾生有涯知无涯
   
3. ![design](https://docs.microsoft.com/en-us/previous-versions/msp-n-p/images/dn589784.ae4c3e59526d69403f5bacc7840b1fb5(en-us,pandp.10).png)

4. ![cn](https://static001.geekbang.org/infoq/2b/2b11bb362cfa71313d8731b422134545.png)

5. 熔断器implement

```go

package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/breaker"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stat"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/internal/response"
)

const breakerSeparator = "://"

// BreakerHandler returns a break circuit middleware.
func BreakerHandler(method, path string, metrics *stat.Metrics) func(http.Handler) http.Handler {
	brk := breaker.NewBreaker(breaker.WithName(strings.Join([]string{method, path}, breakerSeparator)))
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			promise, err := brk.Allow()
			if err != nil {
				metrics.AddDrop()
				logx.Errorf("[http] dropped, %s - %s - %s",
					r.RequestURI, httpx.GetRemoteAddr(r), r.UserAgent())
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}

			cw := &response.WithCodeResponseWriter{Writer: w}
			defer func() {
				if cw.Code < http.StatusInternalServerError {
					promise.Accept()
				} else {
					promise.Reject(fmt.Sprintf("%d %s", cw.Code, http.StatusText(cw.Code)))
				}
			}()
			next.ServeHTTP(cw, r)
		})
	}
}

```


[Circuit Breaker Pattern](https://docs.microsoft.com/en-us/previous-versions/msp-n-p/dn589784(v=pandp.10))

[sony github](https://github.com/sony/gobreaker)

[alibaba](https://github.com/alibaba/sentinel-golang)

[juejin](https://juejin.cn/post/7030997067560386590)