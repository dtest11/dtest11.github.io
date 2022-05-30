package bit

import (
	"fmt"
	"testing"
)

/**
encode[i]=arr[i] XOR arr[i+1]
2个数XOR之后的结果有如下的定理
c = a ^ b
b = c ^ a
a = c ^ b
*/

func decode(encode []int, first int) []int {
	var result = make([]int, len(encode)+1)
	result[0] = first
	prev := first
	for i := 1; i < len(encode)+1; i++ {
		result[i] = prev ^ encode[i-1]
		prev = result[i]
	}
	return result
}

func Test_encode(t *testing.T) {
	a := 10
	b := 2
	c := a ^ b
	fmt.Println(c ^ b)
	fmt.Println(a ^ a)
	d := []int{1, 2, 3}
	fmt.Println(decode(d, 1))

}
