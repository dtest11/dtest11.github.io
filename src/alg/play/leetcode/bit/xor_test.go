package bit

import (
	"fmt"
	"testing"
)

// 这里有个技巧 a ^ a =a
func xorOperation(n int, start int) int {
	prev := start
	for i := 1; i < n; i++ {
		current := start + 2*i
		result := prev ^ current
		prev = result
	}
	fmt.Println(prev)
	return prev
}

// xor 2个不一致才返回1
func Test_xoreOperation(t *testing.T) {
	xorOperation(5, 0)
	xorOperation(4, 3)
}
