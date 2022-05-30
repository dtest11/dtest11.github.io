package tree

/***

Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees
of every node never differ by more than 1.

Example:

Given the sorted array: [-10,-3,0,5,9],

One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

      0
     / \
   -3   9
   /   /
 -10  5

 ***/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/***
这个题目首要 BST 有特性： letf < root < right, 因此可以从中间的节点开始 切割
**/

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid+1:]
	root := &TreeNode{Val: nums[mid]}

	root.Left = sortedArrayToBST(left)
	root.Right = sortedArrayToBST(right)

	return root
}
