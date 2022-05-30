package visit

import (
	"log"

	"github.com/9999-dollor/algorithm/leetcode/stack"
	"github.com/9999-dollor/algorithm/leetcode/tree"
)

// PostOrderRecurse 后序遍历 left-right - root
func PostOrderRecurse(root *tree.TreeNode) {
	if root == nil {
		return
	}
	PostOrderRecurse(root.Left)
	PostOrderRecurse(root.Right)
	log.Printf("%v", root.Val)
}

// PostOrderIterative 非递归的后序遍历 stack, left-right-root
/**
    后序遍历的非递归实现是三种遍历方式中最难的一种。因为在后序遍历中，
	要保证左孩子和右孩子都已被访问并且左孩子在右孩子前访问才能访问根结点，这就为流程的控制带来了难题。下面介绍两种思路。

	要保证根结点在左孩子和右孩子访问之后才能访问，因此对于任一结点P，先将其入栈。如果P不存在左孩子和右孩子，
	则可以直接访问它；或者P存在左孩子或者右孩子，但是其左孩子和右孩子都已被访问过了，则同样可以直接访问该结点。
	若非上述两种情况，则将P的右孩子和左孩子依次入栈，
	这样就保证了每次取栈顶元素的时候，左孩子在右孩子前面被访问，左孩子和右孩子都在根结点前面被访问。
**/
func PostOrderIterative(root *tree.TreeNode) {
	if root == nil {
		return
	}
	stk := stack.NewStack()
	stk.Push(root)
	var prev *tree.TreeNode // 前一次访问的节点
	for !stk.Empty() {
		current := stk.Peek().(*tree.TreeNode)
		if current.Left == nil && current.Right == nil {
			// 如果当前结点没有孩子结点
			prev = current
			stk.Pop()
		}
		if prev != nil && (prev == current.Left || prev == current.Right) {
			// 当前节点的孩子节点都已被访问过
			prev = current
			stk.Pop()
		}
		if current.Right != nil {
			stk.Push(current.Right)
		}
		if current.Left != nil {
			stk.Push(current.Left)
		}
	}
}
