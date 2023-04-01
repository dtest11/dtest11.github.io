package dynamic_programming

import "math"

// 只允许进行一次交易
func maxProfit_K_1(prices []int) int {
	var dp = make([][]int, len(prices))
	for i := range prices {
		dp[i] = make([]int, 2)
	}
	for i, v := range prices {
		if i-1 == -1 {
			/**
			dp[0][0]=Max(dp[0][0],dp[-1][1]+price[0])
			        =max(0, -infinity + prices[i]) = 0
			dp[0][1]=Max(dp[-1][1],dp[-1][0]-price[0],)=-pirce[0]
			*/
			dp[0][0] = 0
			dp[0][1] = -v
		} else {
			dp[i][0] = MaxInt(dp[i-1][0], dp[i-1][1]+v)
			dp[i][1] = MaxInt(dp[i-1][1], -v)
		}

	}
	return dp[len(prices)-1][0]
}

// maxProfit_K_NoLimit k次数无限
func maxProfit_K_NoLimit(prices []int) int {
	l := len(prices)
	var dp = make([][]int, l)
	for i := range prices {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < l; i++ {
		if i == 0 {
			/**
			dp[0][0]=Max(dp[0][0],dp[-1][1]+price[0])
			        =max(0, -infinity + prices[i]) = 0
			dp[0][1]=Max(dp[-1][1],dp[-1][0]-price[0],)=-pirce[0]

			*/
			dp[0][0] = 0
			dp[0][1] = -prices[i]
		} else {
			dp[i][0] = MaxInt(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = MaxInt(dp[i-1][1], dp[i-1][0]-prices[i])
		}

	}
	return dp[l-1][0]
}

// 只能进行2次交易
// k=2
func maxProfit_k_2(prices []int) int {
	var dp = make([][][]int, len(prices))
	for i := range prices {
		dp[i] = make([][]int, 3)
		for j := 0; j <= 2; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < len(prices); i++ {
		//for k := 1; k < 2; k++ { }
		if i == 0 {
			//dp[i][1][0] = MaxInt(dp[-1][1][0], dp[-1][1][1]+prices[i])
			dp[i][1][0] = 0
			//dp[i][1][1] = MaxInt(dp[-1][1][1], -prices[i])
			dp[i][1][1] = -prices[i]
			//dp[i][2][0] = MaxInt(dp[-1][2][0], dp[-1][2][1]+prices[i])
			dp[i][2][0] = 0
			//dp[i][2][1] = MaxInt(dp[-1][2][1], dp[-1][1][0]-prices[i])
			dp[i][2][1] = -prices[i]

			continue
		}
		dp[i][1][0] = MaxInt(dp[i-1][1][0], dp[i-1][1][1]+prices[i])
		dp[i][1][1] = MaxInt(dp[i-1][1][1], -prices[i])

		dp[i][2][0] = MaxInt(dp[i-1][2][0], dp[i-1][2][1]+prices[i])
		dp[i][2][1] = MaxInt(dp[i-1][2][1], dp[i-1][1][0]-prices[i])

	}
	return dp[len(prices)-1][2][0]
}

// maxProfit_N 没有上限
func maxProfit_N(maxK int, prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}
	if maxK > l/2 {
		// 一次交易由买入和卖出构成，至少需要两天。所以说有效的限制 k 应该不超过 n/2，如果超过，就没有约束作用
		return maxProfit_K_NoLimit(prices)
	}
	dp := make([][][]int, l)
	for i := range dp {
		dp[i] = make([][]int, maxK+1)
		for j := 0; j < maxK+1; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	// k== 0
	for i := 0; i < l; i++ {
		dp[i][0][0] = 0
		dp[i][0][1] = math.MinInt
	}
	// 枚举状态
	for i := 0; i < l; i++ {
		for k := 1; k <= maxK; k++ {
			if i == 0 {
				//dp[i][k][0] = MaxInt(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
				dp[i][k][0] = 0
				//dp[i][k][1] = MaxInt(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = MaxInt(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = MaxInt(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[l-1][maxK][0]
}

func maxProfitWithCoolDown(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < l; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
		} else if i == 1 {
			dp[i][0] = MaxInt(dp[0][0], dp[0][1]+prices[i])
			// dp[i][1] = MaxInt(dp[0][1], dp[-1][0]-prices[i])
			dp[i][1] = MaxInt(dp[0][1], -prices[i])

		} else {
			dp[i][0] = MaxInt(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = MaxInt(dp[i-1][1], dp[i-2][0]-prices[i])
		}
	}
	return dp[l-1][0]
}

func maxProfit_with_fee(prices []int, fee int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < l; i++ {
		if i == 0 {
			// dp[0][0] = MaxInt(dp[-1][0], dp[-1][1]+prices[0])
			dp[0][0] = 0
			dp[0][1] = -prices[i] - fee

		} else {
			dp[i][0] = MaxInt(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = MaxInt(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
		}
	}
	return dp[l-1][0]
}


