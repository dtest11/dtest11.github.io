### [ 300.最长递增子序列长度](https://leetcode.cn/problems/longest-increasing-subsequence/)



```text
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7]
的子序列。

输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

```

状态定义:dp[i] 代表0..i 个元素，以nums[i]结尾的最长递增子序列
```go

for i in nums {
   for j :=0;j<i;j++ {
	   dp[i]=1
	   if nums[i]>nums[j]{
		   dp[i]=max(dp[i],dp[j]+1)


// dp[i]=max(dp[j])+1 => nums[i]>nums[j]

```

```go
{{#include longest-increasing-subsequence.go:0:29}}
```

#### 二分查找实现:
原理： 将数据比作扑克牌，最后的目的是要将这些牌分成几个堆，对于拿到的一张牌，从最左边的堆开始查找放置的位置，堆最顶部的牌必须要比它小，才能放置，如果没有找到合适的堆，那么就新建一个堆，将该牌放到这个新堆里面。

```go
{{#include longest-increasing-subsequence.go:31:58}}
```

* 673. [最长递增子序列的个数](https://leetcode.cn/problems/number-of-longest-increasing-subsequence/)
给定一个未排序的整数数组 nums ， 返回最长递增子序列的个数 。

注意 这个数列必须是 严格 递增的。(不包含相等)
[返回最长子序列](https://www.nowcoder.com/questionTerminal/9cf027bf54714ad889d4f30ff0ae5481)

[题解没有看懂TODO](https://blog.nowcoder.net/n/78d57989c2104ebe823368cc1301113f?f=comment)
```text
输入：数组arr; n = |arr| <= 10^5
输出：数组ans，所有最长上升子序列中字典序最小的一个
```

```go
{{#include longest-increasing-subsequence.go:61:85}}
```



