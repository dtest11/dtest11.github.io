package tree

// easy

func isBalanced(root *TreeNode) bool {
	return check(root) != -1
}

func check(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := check(root.Left)
	right := check(root.Right)
	if left == -1 || right == -1 || abs(left-right) > 1 {
		return -1
	}
	return 1 + max(left, right)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
