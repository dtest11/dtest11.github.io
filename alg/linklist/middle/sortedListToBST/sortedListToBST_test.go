package sortedListToBST

import (
	"testing"
)

func Test_buildTree(t *testing.T) {
	var result = []int{-10, -3, 0, 5, 9}
	var tmp = buildTree(result)
	t.Log(tmp.Val)
}
