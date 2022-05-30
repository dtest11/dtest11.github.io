package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	data := [][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}
	output := [][]int{{1, 6}, {8, 10}, {15, 18}}
	result := Merge(data)
	assert.Equal(t, output, result, "合并区间")
	data= [][]int{}
	Merge(data)
}
