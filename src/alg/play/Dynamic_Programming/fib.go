package dp

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fib2(N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	return fib2(N-1) + fib2(N-2)
}

/**
状态转移方程
**/
