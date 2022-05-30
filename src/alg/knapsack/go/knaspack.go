package kanspack

import (
	"golang.org/x/exp/constraints"
)

// 背包问题
// 动态规划问题
// 1. 先确定2个问题，状态和选择
// 0.1 背包问题

type Payload struct {
	Value  int
	Weight int
}

// K01 : dp[i][w]: 定义为前i个物品，在容量为w的情况下的最大值
// 选择： 遍历到第i个商品，(1)将该商品装入背包，(2)不装入背包
// (1): dp[i][w]=dp[i-1][w-data[i].Value]+data[i].Value
// (2): dp[i][w]=dp[i-1][w]
func K01(data []Payload, maxWeight int) int {
	var dp [][]int = make([][]int, len(data))
	for i := 0; i < len(data); i++ {
		dp[i] = make([]int, maxWeight+1)
	}
	// base case
	for i := 1; i < len(data); i++ {
		for w := 1; w <= maxWeight; w++ {
			//dp[i][w] = Max(dp[i-1][w],
			//	dp[i-1][w-data[i].Weight]+data[i].Value,
			//) // 这里需要检查索引 w-data[i].Weight 是否是正数
			curWeight := data[i].Weight
			curVal := data[i].Value
			if w-curWeight > 0 {
				dp[i][w] = Max(dp[i-1][w], dp[i-1][w-curWeight]+curVal)
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}
	return dp[len(data)-1][maxWeight]
}

// KUnlimited  : dp[i][w]: 定义为前i个物品，在容量为w的情况下的最大值
// 选择： 遍历到第i个商品，(1)将该商品装入背包，(2)不装入背包
// (1): dp[i][w]=dp[i][w-data[i].Value]+data[i].Value
// (2): dp[i][w]=dp[i-1][w]
func KUnlimited(data []Payload, maxWeight int) int {
	var dp [][]int = make([][]int, len(data))
	for i := 0; i < len(data); i++ {
		dp[i] = make([]int, maxWeight+1)
	}
	// base case
	for i := 1; i < len(data); i++ {
		for w := 1; w <= maxWeight; w++ {
			//dp[i][w] = Max(dp[i-1][w],
			//	dp[i-1][w-data[i].Weight]+data[i].Value,
			//) // 这里需要检查索引 w-data[i].Weight 是否是正数
			curWeight := data[i].Weight
			curVal := data[i].Value
			if w-curWeight > 0 {
				dp[i][w] = Max(dp[i-1][w], dp[i][w-curWeight]+curVal)
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	return dp[len(data)-1][maxWeight]
}

func Max[T constraints.Ordered](a T, b T) T {
	if a > b {
		return a
	}
	return b
}
