# 买卖股票的最佳时机

| LeetCode | 力扣 | 难度 |
| :----: | :----: | :----: |
| [121. Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | [121. 买卖股票的最佳时机](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/) | 🟢
| [122. Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/) | [122. 买卖股票的最佳时机 II](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/) | 🟠
| [123. Best Time to Buy and Sell Stock III](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/) | [123. 买卖股票的最佳时机 III](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/) | 🔴
| [188. Best Time to Buy and Sell Stock IV](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/) | [188. 买卖股票的最佳时机 IV](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/) | 🔴
| [309. Best Time to Buy and Sell Stock with Cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/) | [309. 最佳买卖股票时机含冷冻期](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/) | 🟠
| [714. Best Time to Buy and Sell Stock with Transaction Fee](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/) | [714. 买卖股票的最佳时机含手续费](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/) | 🟠
| - | [剑指 Offer 63. 股票的最大利润](https://leetcode.cn/problems/gu-piao-de-zui-da-li-run-lcof/) | 🟠


```text
给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

```

状态有那些?

1. 天数(第几天)：
2. 当前的最大的交易次数:k[1...n]
3. 是否持有股票:s[0,1]

dp[i][k][s]
代表第i天，最大交易次数为K，s=0 无股票，s=1,有股票


### [121. 买卖股票的最佳时机](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)

**只允许进行一次交易**

状态转移方程:

```text
对于每一天的选择有，买入,卖出，不动, 而卖出必须要在买入之后
dp[i][k][0]               不包含股票=
dp[i-1][k][1]+price[i]              {昨天持有，今天卖出}
dp[i-1][k][0]                       {昨天没有，今天也没有：休息}


dp[i][k][1]             今天有股票=
dp[i-1][k-1][0]-price[i]            {昨天没有，今天买入}
dp[i-1][k][1]  						{昨天有，今天没有买，休息}
```

针对基本case 需要赋值下
```text
dp[-1][...][0] = 0
解释：因为 i 是从 0 开始的，所以 i = -1 意味着还没有开始，这时候的利润当然是 0。

dp[-1][...][1] = -infinity
解释：还没开始的时候，是不可能持有股票的。
因为我们的算法要求一个最大值，所以初始值设为一个最小值，方便取最大值。

dp[...][0][0] = 0
解释：因为 k 是从 1 开始的，所以 k = 0 意味着根本不允许交易，这时候利润当然是 0。

dp[...][0][1] = -infinity
解释：不允许交易的情况下，是不可能持有股票的。
因为我们的算法要求一个最大值，所以初始值设为一个最小值，方便取最大值。

```


把上面的状态转移方程总结一下：

```text
base case：
dp[-1][...][0] = dp[...][0][0] = 0
dp[-1][...][1] = dp[...][0][1] = -infinity

状态转移方程：
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
```


_针对只允许一次交易的，那么k=1_

状态转移方程:
```text

dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])

dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0] - prices[i]) 
            = max(dp[i-1][1][1], -prices[i])
解释：k = 0 的 base case，所以 dp[i-1][0][0] = 0。

现在发现 k 都是 1，不会改变，即 k 对状态转移已经没有影响了。
可以进行进一步化简去掉所有 k：
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], -prices[i])

```

```go
func maxProfit(prices []int) int {
	var dp = make([][]int, len(prices))
	for i, v := range prices {
		dp[i][0] = MaxInt(dp[i-1][0], dp[i][1]+v)// 注意这个i=0 会有dp[-1][0]
		dp[i][1] = MaxInt(dp[i-1][1], -v)//dp[-1][1]
	}
	return dp[len(prices)][0]
}

```

```go
{{#include best-time-to-buy-and-sell-stock.go:0:27}}
```

### [122 题「买卖股票的最佳时机 II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/)


You are given an integer array prices where prices[i] is the price of a given stock on the ith day.

On each day, you may decide to buy and/or sell the stock. You can only hold at 
most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.

Find and return the maximum profit you can achieve.


**也就是 k 为正无穷的情况**

如果 k 为正无穷，那么就可以认为 k 和 k - 1 是一样的。可以这样改写框架
```text
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
            = max(dp[i-1][k][1], dp[i-1][k][0] - prices[i])

我们发现数组中的 k 已经不会改变了，也就是说不需要记录 k 这个状态了：
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
```
```go
{{#include best-time-to-buy-and-sell-stock.go:29:54}}
```

### [123.买卖股票的最佳时机 III](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/)

You are given an array prices where prices[i] is the price of a given stock on the ith day.

Find the maximum profit you can achieve. You may complete at most two transactions.

Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

**你最多可以完成两笔交易:k<=2**

动态转移方程
由于k的最大值为2，那么状态可能0,1,2,需要将以上的转移方程都罗列出来
k==0 ，不需要统计，因为不做交易数据统计无效

```text            
k=1
dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0] - prices[i])
            = max(dp[i-1][1][1],  - prices[i])

k=2
 dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1] + prices[i])
 dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0] - prices[i])
           
```


```go
{{#include best-time-to-buy-and-sell-stock.go:54:90}}
```

### [188.Best Time to Buy and Sell Stock IV](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/description/)


** 你最多可以完成 k 笔交易:0 <= k <= 100**

You are given an integer array prices where prices[i] is the price of a given stock on the ith day, and an integer k.

Find the maximum profit you can achieve. You may complete at most k transactions.

Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).




1. 一次交易由买入和卖出构成，至少需要两天。所以说有效的限制 k 应该不超过 n/2，如果超过，就没有约束作用了，相当于 k 没有限制的情况，而这种情况是之前解决过的。

```go
{{#include best-time-to-buy-and-sell-stock.go:89:126}}
```

### [309. Best Time to Buy and Sell Stock with Cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)


You are given an array prices where prices[i] is the price of a given stock on the ith day.

Find the maximum profit you can achieve. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times) with the following restrictions:

After you sell your stock, you cannot buy stock on the next day (i.e., cooldown one day).
Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

** k 无穷，但是有交易冷冻期**
```text
dp[i][0]=max(dp[i-1][0],dp[i-1][1]+price[i])
dp[i][1]=max(dp[i-1][1],dp[i-2][0]-pirce[i])// 每次 sell 之后要等一天才能继续交易
```

```go
{{#include best-time-to-buy-and-sell-stock.go:128:154}}
```

### [714. Best Time to Buy and Sell Stock with Transaction Fee](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)

You are given an array prices where prices[i] is the price of a given stock on the ith day, and an integer fee representing a transaction fee.

Find the maximum profit you can achieve. You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction.

Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

每次交易要支付手续费，只要把手续费从利润中减去即可，改写方程：
状态转移方程
```text

dp[i][0]=Max(dp[i-1][0],dp[i-1][1]+price[i])
  
dp[i][1]=Max(dp[i-1][1],dp[i-1]+price[i]-fee)
```

```go
{{#include best-time-to-buy-and-sell-stock.go:155:177}}
```

### [剑指 Offer 63. 股票的最大利润](https://leetcode.cn/problems/gu-piao-de-zui-da-li-run-lcof/)

假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？
K=1 的解法
