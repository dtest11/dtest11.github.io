package dynamic_programming

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fib(t *testing.T) {
	n := 10
	t.Log(fib(n))
	t.Log(fib2(n))
}

func Test_maxSubArray(t *testing.T) {
	var a = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	t.Log(maxSubArray(a))
}

func Test_longest_increasing_subsequence(t *testing.T) {
	var nums = []int{10, 9, 2, 5, 3, 7, 101, 18}
	result := lengthOfLIS(nums)
	t.Log(result)
}

func Test_longestCommonSubsequence(t *testing.T) {
	text1 := "abcde"
	text2 := "ace"
	res := longestCommonSubsequence(text1, text2)
	assert.Equal(t, 3, res)
}

func Test_lengthOfLIS_binary_search(t *testing.T) {
	var nums = []int{10, 9, 2, 5, 3, 7, 101, 18}
	res := lengthOfLISV2(nums)
	assert.Equal(t, 4, res)

	res = lengthOfLIS_binary_search(nums)
	assert.Equal(t, 4, res)
}

func Test_maxProfit_k_Two(t *testing.T) {
	t.Log(maxProfit_N(2, []int{3, 3, 5, 0, 0, 3, 1, 4}))
}

func Test_rob(t *testing.T) {
	var nums = []int{1, 2, 3, 1}
	t.Log(rob(nums))
	t.Log(rob([]int{2, 1, 1, 2}))
	t.Log(robCircle([]int{1, 2, 3, 1}))
	t.Log(robCircle([]int{2, 3, 2}))
	t.Log(robCircle([]int{1, 2, 3}))
}
