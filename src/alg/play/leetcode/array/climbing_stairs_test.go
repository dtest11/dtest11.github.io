package array

import (
	"fmt"
	"testing"
)

func climbStairsRecurse(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairsRecurse(n-1) + climbStairsRecurse(n-2)
}

func climbStairs(n int) int {
	// if n == 0 {
	// 	return 0
	// }
	// if n == 1 {
	// 	return 1
	// }
	// if n == 2 {
	// 	return 2
	// }
	if n < 3 {
		return n
	}
	var prev = 1
	var current = 2
	for i := 3; i <= n; i++ {
		tmp := current
		current = current + prev
		fmt.Println(tmp,current)
		prev =tmp
	}
	return current
}

func Test_climbStairs(t *testing.T) {
	fmt.Println(climbStairs(5))
}
