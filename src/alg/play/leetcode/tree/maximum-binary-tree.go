package tree

// medium
/***

Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:

The root is the maximum number in the array.
The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
Construct the maximum tree by the given array and output the root node of this tree.

Example 1:
Input: [3,2,1,6,0,5]
Output: return the tree root node representing the following tree:

      6
    /   \
   3     5
    \    /
     2  0
       \
        1



**/

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max := nums[0]
	index := 0
	for idx, v := range nums {
		if v >= max {
			max = v
			index = idx
		}
	}

	node := TreeNode{Val: max}
	node.Left = constructMaximumBinaryTree(nums[:index])
	node.Right = constructMaximumBinaryTree(nums[index+1:])
	return &node
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	val := preorder[0]
	root := &TreeNode{Val: val}
	if val > preorder[0] {
		root.Left = bstFromPreorder(preorder[1:])
	} else {
		root.Right = bstFromPreorder(preorder[1:])
	}
	return root
}
