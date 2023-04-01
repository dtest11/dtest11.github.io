package tree

// 递归统计深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 使用层次遍历的方式统计深度
func maxDepthWithLevelTravse(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	queue = append(queue, root) // 注意：这里使用队列
	var result = 0
	for len(queue) != 0 {
		temp := queue
		queue = nil
		for _, v := range temp {
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		result++
	}
	return result + 1
}
