package tree

// medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumEvenGrandparent(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0

	if root.Val&1 == 0 {
		if root.Left != nil && root.Left.Left != nil {
			res += root.Left.Left.Val
		}

		if root.Left != nil && root.Left.Right != nil {
			res += root.Left.Right.Val
		}

		if root.Right != nil && root.Right.Left != nil {
			res += root.Right.Left.Val
		}

		if root.Right != nil && root.Right.Right != nil {
			res += root.Right.Right.Val
		}
	}

	res += sumEvenGrandparent(root.Left)
	res += sumEvenGrandparent(root.Right)
	return res
}
