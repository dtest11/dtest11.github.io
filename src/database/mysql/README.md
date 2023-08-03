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

## why B+
* [数组] 新增数据会导致数据大量移动，不适合
* [二叉树] left<root<right, 如果插入的一直是大于right的数据会导致树的高度一直增长，查询数据时间复杂度为n
* [AVL] (平衡二叉树，每个节点的左子树和右子树的高度差不能超过 1)：不管平衡二叉查找树还是红黑树，都会随着插入的元素增多，而导致树的高度变高，这就意味着磁盘 I/O 操作次数多，会影响整体数据查询的效率。
* B+ 树的插入和删除效率更高：B+的删除和插入都是针对一条路径，因此删除和插入改变的数据很少，更高效
* B+ 非叶子节点只存储了索引项，因此少量的页就能存储大量的索引数据，定位数据更快，读取磁盘次数少
* B+ 所有叶子节点 使用了双向链表，关联，因此支持范围查询，更高效

## Lock
```text
select * from performance_schema.data_locks\G;
```


*  [层高计算](https://cloud.tencent.com/developer/article/2121055)
* 索引下推：联合索引




pt-online-schema-change -uroot -p 123456 --execute --alter "MODIFY COLUMN person_name varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名称';" D=hello,t=person 


select * from person where id = 1 for update;
 

 # insert into 

 ```sql
 insert into T(id,name) 
   select id,name from T2

```

```sql
insert into T(id,name) 
   select 1,3 from dual 
      where not exists(select * from T3);
```

# create index
```sql
create unique index  uniq_idx_firstname on actor (first_name);
create index idx_lastname on actor (last_name);
```

# create view
```sql
CREATE VIEW view_name AS
SELECT column1, column2, ...
FROM table_name
WHERE condition;
```

# index
```sql
select *
from salaries
force index (idx_emp_no)
where emp_no=10005
```

# alter
```sql
alter table actor add create_date datetime not null default '2020-10-01 00:00:00';

CREATE TRIGGER trig1 AFTER INSERT
ON employees_test FOR EACH ROW
 INSERT INTO audit VALUES(new.id,new.name);


UPDATE titles_test
SET emp_no = REPLACE(emp_no,
                     10001,
                     10005)
where id = 5;

rename table A to B;


alter table audit add constraint foreign key (emp_no) references  employees_test(id)

//返回部分字段
LEFT(s,n)返回字符串 s 的前 n 个字符
RIGHT(s,n)返回字符串 s 的后 n 个字符
第一种：（right）

select first_name
from employees
order by right(first_name,2)
第二种：（substr）
select first_name from employees  order by substr(first_name,-2)

第三种：（substring）

select first_name from employees  order by substring(first_name,-2)


group_concat（）函数将group by产生的同一个分组中的值连接起来，返回一个字符串结果

select dept_no, group_concat(emp_no) as employees
from dept_emp
group by dept_no


SELECT (SUM(salary) - MAX(salary) - MIN(salary)) / (COUNT(1) - 2) as avg_salary
FROM salaries
where to_date = '9999-01-01';

#分页查询employees表，每5行一页，返回第2页的数据
LIMIT 语句结构： LIMIT X,Y 
Y ：返回几条记录
X：从第几条记录开始返回（第一条记录序号为0，默认为0）
select *
from employees
limit 5 , 5

#窗口函数
select emp_no,
    salary,
    sum(salary) over(
        order by emp_no
    )
from salaries;




```

获取row number
```sql

select emp_no,
    row_number() OVER (ORDER BY id)
from salaries;
```

having
```sql
SELECT number
from grade
GROUP BY number
HAVING(count(number)) >= 3;
```