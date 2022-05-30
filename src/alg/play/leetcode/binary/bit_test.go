package binary

import (
	"fmt"
	"testing"
)

func Test_hello(t *testing.T) {
	fmt.Println(8 >> 2)
	fmt.Println(8 << 2)
	fmt.Println(isOdd(0))
	fmt.Println(isOdd(1))
	fmt.Println(isOdd(2))
	fmt.Println(isOdd(3))

	// & 2位1=1
	// | 1个1=1
	// ^  2个数相同才是1
}

// 奇数最后一位肯定是1 ，所以跟0做个与运算（&），判断留下来的值
func isOdd(n int) bool {
	last := n & 1
	return last == 0
}
