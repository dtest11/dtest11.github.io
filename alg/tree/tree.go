package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rebuildTreeFromLevelRecurse(data []int) *TreeNode {
	panic("not implement")

	idx := 0
	var root TreeNode
	for idx < len(data) {
		val := data[0]
		data = data[0:]
		root.Val = val

		leftVal := data[0]
		data = data[0:]
		rightVal := data[0]
		data = data[0:]

		root.Left = &TreeNode{Val: leftVal}
		root.Right = &TreeNode{Val: rightVal}
	}
	return &root

}

func levelTravse(root *TreeNode) []int {
	queue := []*TreeNode{}
	queue = append(queue, root)
	var result []int
	for len(queue) != 0 {
		root := queue[0]
		queue = queue[0:]
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
		result = append(result, root.Val)
	}
	return result
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	left := inorderTraversal(root.Left)
	result = append(result, left...)
	result = append(result, root.Val)
	right := inorderTraversal(root.Right)
	result = append(result, right...)
	return result
}
