package tree

/**
 * Definition for a Node.

 */

type Node struct {
	Val      int
	Children []*Node
}

// 深度遍历
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	return maxDepthHelp(root, depth)
}

func maxDepthHelp(root *Node, depth int) int {
	if root == nil {
		return depth
	}
	var queue []*Node
	queue = append(queue, root)
	for len(queue) != 0 {
		var newQue []*Node
		for _, v := range queue {
			sameLevel := v.Children
			newQue = append(newQue, sameLevel...)
		}
		depth += 1
		queue = newQue
	}
	return depth
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isUnivalTreeHelp(root, root.Val)
}

func isUnivalTreeHelp(root *TreeNode, last int) bool {
	if root == nil {
		return true
	}
	if root.Val != last {
		return false
	}
	left := isUnivalTreeHelp(root.Left, last)
	right := isUnivalTreeHelp(root.Right, last)
	return left && right
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
