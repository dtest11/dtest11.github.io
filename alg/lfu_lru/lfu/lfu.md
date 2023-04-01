### LFU


[leetcode](https://leetcode.com/problems/lfu-cache/description/)


* def

***如果一个数据在最近一段时间很少被访问到，那么可以认为在将来它被访问的可能性也很小。因此，当空间满时，最小频率访问的数据最先被淘汰。***

定义一个字典:key 是使用频率，value 是一个链表，链表根据使用时间排序，最新的数据在头部，
历史数据在胃部，容量超过的时候，删除最小的频率下的最后一个数据

. 如果最小的频率就一个数据，需要增加最小频率+1，不然下次驱逐的时候找不到合适的数据

[href](https://www.cnblogs.com/thisiswhy/p/14266381.html)
 
* impl

```go
{{#include ./v1/v1.go:1:25}}
```

* Get

```go
{{#include ./v1/v1.go:55:62}}
```

* Put


```go
{{#include ./v1/v1.go:64:88}}
```

* 统计频率

```go
{{#include ./v1/v1.go:26:53}}
```

*** 双端链表实现
```go
{{#include ./v2/v2.go}}
```