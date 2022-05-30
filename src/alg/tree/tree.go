package tree

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func rebuildTreeFromLevelRecurse(data []int) *TreeNode {
	panic("not implement")
	if len(data) == 0 {
		return nil
	}
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
