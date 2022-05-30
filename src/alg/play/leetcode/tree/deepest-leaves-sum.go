package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }



 */

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}

	var record []*TreeNode
	level := 1
	LastLevel := level
	for len(queue) != 0 {
		cur := queue[:]
		var tmp []*TreeNode
		level++
		for _, e := range cur {
			if e.Left == nil && e.Right == nil {
				if LastLevel == level {
					record = append(record, e)
				} else {
					record = []*TreeNode{e}
					LastLevel = level
				}
			}
			if e.Right != nil {
				tmp = append(tmp, e.Right)
			}
			if e.Left != nil {
				tmp = append(tmp, e.Left)
			}
		}
		queue = tmp
	}

	result := 0
	for _, v := range record {
		result += v.Val
	}
	return result
}
