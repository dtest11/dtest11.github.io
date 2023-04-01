package dynamic_programming

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fib2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	prev := 1
	cur := 1
	for i := 3; i <= n; i++ {
		sum := cur + prev
		prev = cur
		cur = sum
	}
	return cur
}

// 给你k种面值的硬币，面值分别为c1, c2 ... ck，每种硬币的数量无限，再给一个总金额amount，
// 问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1 。算法的函数签名如下：
// dp = min(dp[n-coins]+1)
func coinChange(amount int, coins []int) int {
	/**
	dp[n]=dp[n-1]+1
	*/
	dp := make([]int, amount+1, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 0; i < amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = func() int {
				if dp[i] > dp[i-coin]+1 {
					return dp[i-coin] + 1
				}
				return dp[i]
			}()
		}
	}
	if dp[amount+1] == amount+1 {
		return -1
	}
	return dp[amount+1]
}

func longeset() {

}
