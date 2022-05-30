package dp

import "testing"

func Test_longest_increasing_subsequence(t *testing.T) {
	var nums = []int{10, 9, 2, 5, 3, 7, 101, 18}
	result := longest_increasing_subsequence(nums)
	t.Log(result)
}
