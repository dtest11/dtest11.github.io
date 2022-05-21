限流

1. 常见限流算法

   

   - ###### 计数器限流算法
     
     我们可以直接通过一个计数器，限制每一秒钟能够接收的请求数。比如说 qps定为 1000，那么实现思路就是从第一个请求进来开始计时，在接下去的 1s 内，每来一个请求，就把计数加 1，如果累加的数字达到了 1000，那么后续的请求就会被全部拒绝。等到 1s 结束后，把计数恢复成 0 ，重新开始计数。
     
     缺点：如果1s 内的前半秒，已经通过了 1000 个请求，那后面的半秒只能请求拒绝，我们把这种现象称为“突刺现象”。
     
     
     
     ```go 
     package main
     
     import (
        "log"
        "sync"
        "time"
     )
     
     type Counter struct {
        rate  int           // The maximum number of requests allowed in the count cycle
        begin time.Time     // Count start time
        cycle time.Duration // Counting cycle
        count int           // Cumulative number of requests received in the counting cycle
        lock  sync.Mutex
     }
     
     func (l *Counter) Allow() bool {
        l.lock.Lock()
        defer l.lock.Unlock()
     
        if l.count == l.rate-1 {
           now := time.Now()
           if now.Sub(l.begin) >= l.cycle {
              //Within the allowable speed range,   Reset counter
              l.Reset(now)
              return true
           } else {
              return false
           }
        } else {
           //Rate limit not reached, count plus 1
           l.count++
           return true
        }
     }
     
     func (l *Counter) Set(r int, cycle time.Duration) {
        l.rate = r
        l.begin = time.Now()
        l.cycle = cycle
        l.count = 0
     }
     
     func (l *Counter) Reset(t time.Time) {
        l.begin = t
        l.count = 0
     }
     
     func main() {
        var wg sync.WaitGroup
        var lr Counter
        lr.Set(3, time.Second) //  Up to 3 requests in 1s
        for i := 0; i < 10; i++ {
           wg.Add(1)
           Log.Println("create request:", i)
           go func(i int) {
              if lr.Allow() {
                 Log.Println("response request:", i)
              }
              wg.Done()
           }(i)
     
           time.Sleep(200 * time.Millisecond)
        }
        wg.Wait()
     }
     ```



   - 滑动时间窗算法

        ![o]https://segmentfault.com/img/remote/1460000039958566
        https://github.com/RussellLuo/slidingwindow

​      

   - 漏桶（Leaky Bucket）

     原理就是一个固定容量的漏桶，按照固定速率流出水滴。用过水龙头都知道，打开龙头开关水就会流下滴到水桶里，而漏桶指的是水桶下面有个漏洞可以出水。如果水龙头开的特别大那么水流速就会过大，这样就可能导致水桶的水满了然后溢出。

   - 令牌桶算法（token Bucket）
     
     我们有一个固定的桶，桶里存放着令牌`（token）`。一开始桶是空的，系统按固定的时间`（rate）`往桶里添加令牌，直到桶里的令牌数满，多余的请求会被丢弃。当请求来的时候，从桶里移除一个令牌，如果桶是空的则拒绝请求或者阻塞。
     
   - 漏筒vs token bucket 
     
     
     ![alt 属性文本](https://static001.geekbang.org/infoq/9f/9f645abbc8477ff790ab663b4d75f21e.png)
     
     
     ![alt 属性文本](https://static001.geekbang.org/infoq/47/47bba937ecc3f81b975e9a59a743b3d4.png)
     [vs](https://m.haicoder.net/note/nginx-interview/nginx-interview-nginx-leaky-bucket.html)
     
     除了要求能够限制数据的平均传输速率外，还要求允许某种程度的突发传输

     > 相比于 TokenBucket 中，只要桶内还有剩余令牌，调用方就可以一直消费的策略。Leaky Bucket 相对来说更加严格，调用方只能严格按照预定的间隔顺序进行消费调用

- Leaky Bucket
```go

// Copyright (c) 2016,2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package ratelimit // Package ratelimit import "go.uber.org/ratelimit"

import (
	"sync"
	"time"
)

type mutexLimiter struct {
	sync.Mutex
	last       time.Time
	sleepFor   time.Duration
	perRequest time.Duration
	maxSlack   time.Duration
	clock      Clock
}

// newMutexBased returns a new atomic based limiter.
func newMutexBased(rate int, opts ...Option) *mutexLimiter {
	// TODO consider moving config building to the implementation
	// independent code.
	config := buildConfig(opts)
	perRequest := config.per / time.Duration(rate)
	l := &mutexLimiter{
		perRequest: perRequest,
		maxSlack:   -1 * time.Duration(config.slack) * perRequest,
		clock:      config.clock,
	}
	return l
}

// Take blocks to ensure that the time spent between multiple
// Take calls is on average time.Second/rate.
func (t *mutexLimiter) Take() time.Time {
	t.Lock()
	defer t.Unlock()

	now := t.clock.Now()

	// If this is our first request, then we allow it.
	if t.last.IsZero() {
		t.last = now
		return t.last
	}

	// sleepFor calculates how much time we should sleep based on
	// the perRequest budget and how long the last request took.
	// Since the request may take longer than the budget, this number
	// can get negative, and is summed across requests.
	t.sleepFor += t.perRequest - now.Sub(t.last)

	// We shouldn't allow sleepFor to get too negative, since it would mean that
	// a service that slowed down a lot for a short period of time would get
	// a much higher RPS following that.
	if t.sleepFor < t.maxSlack {
		t.sleepFor = t.maxSlack
	}

	// If sleepFor is positive, then we should sleep now.
	if t.sleepFor > 0 {
		t.clock.Sleep(t.sleepFor)
		t.last = now.Add(t.sleepFor)
		t.sleepFor = 0
	} else {
		t.last = now
	}

	return t.last
}

```
- token bucket
[https://pandaychen.github.io/2020/04/05/GOLANG-X-RATELIMIT-ANALYSIS/](https://pandaychen.github.io/2020/04/05/GOLANG-X-RATELIMIT-ANALYSIS/)

- 自适应限流

思路： 只要我们的 CPU 负载超过 80% 的时候，获取过去 5s 的最大吞吐数据，然后再统计当前系统中的请求数量，只要当前系统中的请求数大于最大吞吐那么我们就丢弃这个请求。

[https://juejin.cn/post/6965406931066290206](https://juejin.cn/post/6965406931066290206)