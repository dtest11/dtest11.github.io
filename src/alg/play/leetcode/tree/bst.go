package tree

import (
	"fmt"
)

func increasingBST(root *TreeNode) *TreeNode {
	var dummy *TreeNode
	return increasingBSTHelp(root, dummy)
}

func increasingBSTHelp(root *TreeNode, dummy *TreeNode) *TreeNode {
	if root == nil {
		return dummy
	}
	left := increasingBSTHelp(root.Left, root)
	root.Left = nil
	root.Right = increasingBSTHelp(root, dummy)
	return left
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root, _ := build(preorder, inorder)
	return root
}

func build(preorder []int, inorder []int) (*TreeNode, []int) {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil, preorder
	}
	node := &TreeNode{Val: preorder[0]}
	var left []int
	var right []int
	var index int
	for ; index < len(inorder); index++ {
		if inorder[index] != preorder[0] {
			left = append(left, inorder[index])
		} else {
			break
		}
	}
	fmt.Println(index, len(inorder))
	if index+1 <= len(inorder) {
		right = inorder[index+1:]
	}
	preorder = preorder[1:]
	if len(left) != 0 {
		node.Left, preorder = build(preorder, left)
	}
	if len(right) != 0 {
		node.Right, preorder = build(preorder, right)
	}
	return node, preorder
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}

	var records []float64
	for len(stack) != 0 {
		var sum int
		var children []*TreeNode
		var nodes int
		for _, i2 := range stack {
			sum += i2.Val
			if i2.Right != nil {
				children = append(children, i2.Right)
			}

			if i2.Left != nil {
				children = append(children, i2.Left)
			}
		}
		records = append(records, float64(sum)/float64(nodes))
		stack = children
	}
	return records
}
