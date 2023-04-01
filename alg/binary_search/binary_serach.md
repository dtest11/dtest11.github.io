* 300.  最长递增子序列长度
[Link text Here](https://leetcode.cn/problems/longest-increasing-subsequence/)



```go
{{#include binary_serach.go:0:23}}
```

* 搜索左侧边界
```go
{{#include binary_serach.go:24:40}}

```
* 查找右边界
```go
{{#include binary_serach.go:42:59}}

```
* 875.爱吃香蕉的珂珂

求吃香蕉的最慢速度，根据题目速度应该在1.piles[n-1],
然后求left_bound
```
珂珂喜欢吃香蕉。这里有 n 堆香蕉，第 i 堆中有piles[i]根香蕉。警卫已经离开了，将在 h 小时后回来。

珂珂可以决定她吃香蕉的速度 k （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 k 根。如果这堆香蕉少于 k 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。

珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。

返回她可以在 h 小时内吃掉所有香蕉的最小速度 k（k 为整数）。



输入：piles = [3,6,7,11], h = 8
输出：4
```

```go
{{#include binary_serach.go:58:85}}
```

* 1011 在 D 天内送达包裹的能力

[href](https://leetcode.cn/problems/capacity-to-ship-packages-within-d-days/description/)

传送带上的包裹必须在 days 天内从一个港口运送到另一个港口。

传送带上的第 i 个包裹的重量为 weights[i]。每一天，我们都会按给出重量（weights）的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。

返回能在 days 天内将传送带上的所有包裹送达的船的最低运载能力。

 
```text
示例 1：

输入：weights = [1,2,3,4,5,6,7,8,9,10], days = 5
输出：15
解释：
船舶最低载重 15 就能够在 5 天内送达所有包裹，如下所示：
第 1 天：1, 2, 3, 4, 5
第 2 天：6, 7
第 3 天：8
第 4 天：9
第 5 天：10

请注意，货物必须按照给定的顺序装运，因此使用载重能力为 14 的船舶并将包装分成 (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) 是不允许的。

```

```go
{{#include binary_serach.go:88:135}}
```

* . 分巧克力 TODO 
```go
{{#include binary_serach.go:137:156}}
```