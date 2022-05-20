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





[https://juejin.cn/post/6965406931066290206](https://juejin.cn/post/6965406931066290206)