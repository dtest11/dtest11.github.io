package pow

func myPowPositiveN(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}

	// 判断n 是否为偶数
	mid := n / 2
	temp := myPowPositiveN(x, mid)
	if n%2 == 0 {
		return temp * temp
	}
	// n 是奇数
	return temp * temp * x
}

// 注意处理n 为负数的情况
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return myPowPositiveN(x, n)
	}
	// 负数： 正数a的-r次幂(r为任何正数)定义为a的r次幂的倒数
	return 1.0 / myPowPositiveN(x, -n)
}

// go

func fastMul(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / fastMul(x, -n)
	}
	if n == 0 {
		return 1
	}
	res := 1.0
	for n > 1 {
		if n%2 != 0 {
			res = res * x
			n = n - 1
		}
		n = n / 2
		x = x * x
	}
	return res * x
}

// isPowerOfTwoV1 试除法
func isPowerOfTwoV1(n int) bool {
	if n < 0 {
		return false
	}
	if n == 0 {
		return true
	}
	for n != 1 {
		if n%2 != 0 {
			return false
		}
		n = n / 2
	}
	return true
}

// 二进制操作
// 一个数

// n 是 2的幂，当且仅当 是正整数，并且 n 的二进制表示中仅包含 1 个 1
// 如果n是2的幂，那么n-1,最高位为0,其余位置为1
func isPowerOfTwoV2(n int) bool {
	if n&(n-1) == 0 {
		return true
	}
	return false
}

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	if n < 0 {
		return false
	}
	for n != 1 {
		if n%3 != 0 {
			return false
		}
		n = n / 2
	}
	return true
}

func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}
	if n < 0 {
		return false
	}
	for n != 1 {
		if n%4 != 0 {
			return false
		}
		n = n / 4
	}
	return true
}
