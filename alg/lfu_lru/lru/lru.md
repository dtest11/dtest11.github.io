### LRU


[leetcode](https://leetcode.com/problems/lfu-cache/description/)


* def

***如果一个数据在很少被访问到，那么可以认为在将来它被访问的可能性也很小。因此，当空间满时，最小频率访问的数据最先被淘汰。***
定义一个链表，根据使用顺序排序， 最新的数据放在头部，历史数据放在头部

> 容量超过的时候驱逐最后的一个数据

[href](https://www.cnblogs.com/thisiswhy/p/14266381.html)

* impl

```go
{{#include ./v2/v2.go:1:15}}
```

* Get

```go
{{#include ./v2/v2.go:17:25}}
```

* Put


```go
{{#include ./v2/v2.go:27:55}}
```


### 双端链表实现

```go
{{#include ./v2/doubly_linklist.go}}
```