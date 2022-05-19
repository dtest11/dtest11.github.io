# redis

* LRU
 * 近最少使用（Least Recently Used，LRU）
 * [go-lru](https://github.com/golang/groupcache/blob/master/lru/lru.go)
    1. 情况一：当有新数据插入时，LRU 算法会把该数据插入到链表头部，同时把原来链表头部的数据及其之后的数据，都向尾部移动一位。
    2. 情况二：当有数据刚被访问了一次之后，LRU 算法就会把该数据从它在链表中的当前位置，移动到链表头部。同时，把从链表头部到它当前位置的其他数据，都向尾部移动一位。
    3. 情况三：当链表长度无法再容纳更多数据时，若再有新数据插入，LRU 算法就会去除链表尾部的数据，这也相当于将数据从缓存中淘汰掉。
 * Redis 中近似 LRU 算法的实现
   1. redisObject上有一个lock 标记访问时间
   2. 判断当前内存使用情况（访问key的时候）（reids.conf maxmemory)
   3. 随机获取一些key 放到待驱逐的数组中（排序）
   4. 选择被淘汰的键值对并删除 (同步或者异步，根据配置 惰性删除)

* LFU
  * 数据访问的频率来选择被淘汰数据的(最不经常使用算法)
  * [go-lfu](https://github.com/dgrijalva/lfu-go/blob/master/lfu.go)


