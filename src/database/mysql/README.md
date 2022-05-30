# mysql


* [B树、B+树索引算法原理（上）](https://www.codedump.info/post/20200609-btree-1/)
* online DDL (copy | inplace, 添加普通索引inplace,不会重建)
  * [MySQL之Online DDL](https://www.modb.pro/db/100527)
  * [DDL是否会索表](http://mysql.taobao.org/monthly/2021/03/06/)
* count
  * 对于 count(主键 id) 来说，InnoDB 引擎会遍历整张表，把每一行的 id 值都取出来，返回给 server 层。server 层拿到 id 后，判断是不可能为空的，就按行累加。
  * 对于 count(1) 来说，InnoDB 引擎遍历整张表，但不取值。server 层对于返回的每一行，放一个数字“1”进去，判断是不可能为空的，按行累加。
  * 对于 count(字段) 来说：
    1. 如果这个“字段”是定义为 not null 的话，一行行地从记录里面读出这个字段，判断不能为 null，按行累加；
    2. 如果这个“字段”定义允许为 null，那么执行的时候，判断到有可能是 null，还要把值取出来再判断一下，不是 null 才累加。
