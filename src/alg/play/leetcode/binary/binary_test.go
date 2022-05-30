package binary

import (
	"fmt"
	"testing"
)

/**
https://www.cnblogs.com/gaizai/p/4233780.html
*/

// Ten2Two 10=>2
func Ten2Two(n int) []int {
	var result []int
	for n > 0 {
		reminder := n % 2 // 余数
		result = append(result, reminder)
		n = n / 2
		fmt.Println(n)
	}
	return result
}

// Two2Ten 2进制转10进制
func Two2Ten([]int) int {
	panic("not implement")
}
func TestTen2Two(t *testing.T) {
	fmt.Println(Ten2Two(43))
}


