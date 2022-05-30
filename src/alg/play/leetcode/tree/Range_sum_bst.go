package tree

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	sum := 0
	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}
	return rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R) + sum
}

// 非递归版本
func IterativerangeSumBST(root *TreeNode, L int, R int) int {
	// func rangeSumBST(root *TreeNode, L int, R int) int {

	if root == nil {
		return 0
	}
	var stack []*TreeNode
	var sum int
	stack = append(stack, root)
	for len(stack) != 0 {
		element := stack[0]
		stack = stack[1:]
		if element != nil {
			if element.Val >= L && element.Val <= R {
				sum += element.Val
			}
			if element.Val > L {
				stack = append(stack, element.Left)
			}
			if element.Val < R {
				stack = append(stack, element.Right)
			}
		}
	}
	return sum
}
