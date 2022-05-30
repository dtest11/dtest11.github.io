package leetcode

import (
	"fmt"
	"math"
)

/****
10 进制数字 翻转

对数字每次除以 10

100 / 10 = 10
10 / 10 = 1

1/ 10

**/
func reverse(x int) int {
	res := 0

	if x == 0 {
		return 0
	}

	for x != 0 {
		fmt.Println("----", x%10, x/10)
		res = 10*res + (x % 10)
		x = x / 10
		fmt.Println(x, res, "fuck")
	}
	fmt.Println("--------------", res)
	if math.MaxInt64 > res || res < math.MinInt64 {
		return 0
	} else {
		return res
	}

}

// func main() {
// 	a := reverse(123)
// 	fmt.Println(a)
// }
