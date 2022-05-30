package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//
//func TestRebuildTree(t *testing.T) {
//	pre := []int{1, 2, 4, 5, 3}
//	ino := []int{4, 2, 5, 1, 3}
//
//	//root := buildTree(pre, ino)
//	//visit.PreOrder(root)
//}

func prepare() *TreeNode {

	/*
				10
			   /  \
			 20	   30
			/ \      \
		   40  50     60
		  /
		 70
	*/

	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	//root.Left.Left = &TreeNode{Val: 40}
	//root.Left.Right = &TreeNode{Val: 50}
	//root.Right.Right = &TreeNode{Val: 60}
	//root.Left.Left.Left = &TreeNode{Val: 70}
	return root
}

func TestTwoSum(t *testing.T) {
	res := findTarget(prepare(), 3)
	assert.Equal(t, true, res, "should equal")

	var a []int = []int{1, 2, 3, 4}
	fmt.Println(a[1:])
	//res = findTarget(prepare(), 1)
	//assert.Equal(t, true, res, "should equal")
}
