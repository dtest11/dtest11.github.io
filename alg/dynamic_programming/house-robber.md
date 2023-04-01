### 打家劫舍



小偷去偷家，有些规则需要遵守，比如相邻的不能抢
#### [198. House Robber](https://leetcode.com/problems/house-robber/)

状态：
1. 屋子顺序
2. 屋子是否被抢了

动态方程:


```text
dp[i][0] :第i间屋子没有被抢 从0..i 的最大值
dp[i][1]: 第i间屋子被抢了
最大的结果：max(最后一间屋子被抢,最后一间屋子没有被抢)
result =Max(dp[n][0],dp[n][1])
当前屋子没有被抢=max(上一个屋子被抢，上一个屋子没有被抢)
dp[i][0]=max(dp[i-1][1],dp[i-1][0])
当前屋子被抢=nums[i]+上一个屋子没有被抢
dp[i][1]=nums[i]+dp[i-1][0]
```

```go
{{#include house-robber.go:0:31}}
```

#### [213.house-robber-ii](https://leetcode.com/problems/house-robber-ii/)

1. 房子链接成一个圆，头和尾不能同时抢了，负责会触发报警
2. 相邻的不能同时抢

题解: 要么第一个不抢，要么最后一个不抢，求这2个情况的max
max(rob[1..n],rob(0..n-1))



```go
{{#include house-robber.go:32:44}}
```


#### [337. House Robber III](https://leetcode.com/problems/house-robber-iii/)

二叉树上打劫

小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。

除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。

给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 

![MarineGEO circle logo](https://assets.leetcode.com/uploads/2021/03/10/rob1-tree.jpg)


```go
{{#include house-robber.go:48:100}}
```