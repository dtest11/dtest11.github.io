* 50. Pow(x, n)

实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）
```text
  n     n/2     n/2
x    = x     * x 

```
根据上面更公式可以先算出一半的值，然后再计算
注意负数的幂
```text
正数a的-r次幂(r为任何正数)定义为a的r次幂的倒数
```

. 递归解法 O(log⁡n)
```go
{{#include pow.go:0:29}}
```
* 快速幂解法

[href](https://www.cnblogs.com/faterazer/p/10978074.html)




```go
{{#include pow.go:32:50}}
```

* 231. isPowerOfTwo

```go
{{#include pow.go:51:78}}
```

* 326. n的三次方 (试除法)
* 342. n的四次方
```go
{{#include pow.go:80:112}}
```