package tree

// easy
// 镜像树
// https://assets.leetcode.com/uploads/2021/02/19/symtree1.jpg
// 1. 层次遍历， 比较每一个层级的元素 是不是一致
// 2. 分开2个树，left ,right tree 比较
// left == right
func isSymmetric(root *TreeNode) bool {
	// base case
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)

}

func helper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	// 继续比较
	return helper(left.Left, right.Right) && helper(left.Right, right.Left)

}
