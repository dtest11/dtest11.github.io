package dfs

import "testing"

func Test_sumRootToLeaf(t *testing.T) {
	//var seed = []int{1, 0, 1, 0, 1, 0, 1}
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 1}

	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 1}

	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 1}
	t.Log(sumRootToLeaf(&root))
	t.Log(maxDepth(&root))
}



