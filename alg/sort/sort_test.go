package sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubble(t *testing.T) {
	data := []int{10, 4, 2, 3, 7, 9, 1}
	Bubble(data)
	t.Log("hello")
	t.Log(data)
}

type Args struct {
	input  []int
	expect []int
}

var args = []Args{
	{
		input:  []int{3, 2, 4, 1},
		expect: []int{1, 2, 3, 4},
	},
	{
		input:  []int{9, 12, 5, 10, 14, 3, 10},
		expect: []int{3, 5, 9, 10, 10, 12, 14},
	},
}

func Test_quick(t *testing.T) {
	for _, arg := range args {
		quick(arg.input, 0, len(arg.input)-1)
		assert.Equal(t, arg.expect, arg.input)
	}
}

func Test_ms(t *testing.T) {
	//var nums = []int{3, 2, 4, 1}
	//nums = mergeSort(nums)
	//fmt.Println(nums)
	nums1 := []int{0, 5, 2, 3, 1, 8}
	fmt.Println(mergeSort(nums1))
}
