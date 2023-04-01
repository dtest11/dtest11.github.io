package array

import "testing"

func Test_merge(t *testing.T) {
	a := []int{1, 2, 3, 0, 0, 0}
	b := []int{2, 5, 6}
	merge(a, 3, b, 3)
	t.Log(a)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
