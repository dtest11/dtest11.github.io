动态规划

1. 最值问题
最长递增子序列 / 最小编辑距离

2. 
都有重复的，最优子结构， dp table/ 备忘录
2.1 状态转移方程
2.2 base case

明确「状态」 -> 定义 dp 数组/函数的含义 -> 明确「选择」-> 明确 base case。

1. 斐波那契数列
dp[n]=dp[n-1]+dp[n-2]

3. 凑零钱问题
给你K中面值的硬币，c1,c2..ck, 总额amount，最少需要多少硬币才能凑出来？
对于dp[amount]=min(dp[amount-n]{n=0...k})+1

4. 最长递增子序列(longest-increasing-subsequence.md)
定义状态: dp[i] 表示0..i的最长递增子序列,结果为max(dp[i])
dp[i+1]=max(dp[i],
dp[i]+1,如果这个条件成立 {str[i]<str[i+]})

5. [最大子序和](dynamic_programming/maximum-subarray.md)
给定一个数组，求一个具有最大和的连续数组

base case dp[0]=nums[0]
```go

dp[i]=max(dp[i-1]+nums[i],nums[i])
```
- [买卖股票的最佳时机](best-time-to-buy-and-sell-stock.md)
- [house-robber](house-robber.md)
