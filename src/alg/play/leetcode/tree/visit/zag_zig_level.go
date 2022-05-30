package visit

// z 字型 层次遍历binary tree
import (
	"github.com/9999-dollor/algorithm/leetcode/stack"
	"github.com/9999-dollor/algorithm/leetcode/tree"
)

func zigzag_level_order(root *tree.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var st stack.Stack
	var reverse = false
	st.Push(root)

	for !st.Empty() {
		l := st.Len()
		var tmp []int

		for i := 0; i < l; i++ {
			node := st.Pop().(*tree.TreeNode)
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				st.Push(node.Left)
			}
			if node.Right != nil {
				st.Push(node.Right)
			}
		}
		if reverse {
			for i, j := 0, len(tmp); i < j; i, j = i+1, j-1 {
				tmp[i], tmp[j] = tmp[j], tmp[j]
			}
			reverse = !reverse
		}
		result = append(result, tmp)
	}
	return result
}
