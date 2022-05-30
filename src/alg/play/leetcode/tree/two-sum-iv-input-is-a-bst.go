package tree

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
滑动窗口，inorder 遍历出来是 升
*/
func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	var result []int
	result = helpInorder(root, result)
	left := 0
	right := len(result) - 1
	//for _, i := range result {
	//	if i == k {
	//		return true
	//	}
	//}
	for left != right {
		start := result[left]
		end := result[right]
		sum := start + end
		if sum > k {
			right--
		}
		if sum == k {
			return true
		}

		if sum < k {
			left += 1
		}
	}
	return false

}

func helpInorder(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}
	result = helpInorder(root.Left, result)
	result = append(result, root.Val)
	result = helpInorder(root.Right, result)
	return result
}
