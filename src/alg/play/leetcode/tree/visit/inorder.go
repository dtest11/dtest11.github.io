package visit

import (
	"fmt"

	"github.com/9999-dollor/algorithm/leetcode/stack"
	"github.com/9999-dollor/algorithm/leetcode/tree"
)

/**
根据中序遍历的顺序，对于任一结点，优先访问其左孩子，而左孩子结点又可以看做一根结点，然后继续访问其左孩子结点，
直到遇到左孩子结点为空的结点才进行访问，然后按相同的规则访问其右子树。因此其处理过程如下：

   对于任一结点P，

  1)若其左孩子不为空，则将P入栈并将P的左孩子置为当前的P，然后对当前结点P再进行相同的处理；

  2)若其左孩子为空，则取栈顶元素并进行出栈操作，访问该栈顶结点，然后将当前的P置为栈顶结点的右孩子；

  3)直到P为NULL并且栈为空则遍历结束

https://play.golang.org/p/coQnnCrBV6

  **/

func Inorder(root *tree.TreeNode) {
	if root == nil {
		return
	}
	var newStack = stack.NewStack()
	cur := root
	for cur != nil || !newStack.Empty() {
		for cur != nil {
			newStack.Push(cur)
			cur = cur.Left
		}
		/* Current must be NULL at this point */
		cur = newStack.Peek().(*tree.TreeNode)
		newStack.Pop()
		fmt.Printf("%d  ", cur.Val)
		cur = cur.Right
	}
}

func InOrderRecursive(root *tree.TreeNode) {
	if root == nil {
		return
	}
	InOrderRecursive(root.Left)
	fmt.Printf("%d  ", root.Val)
	InOrderRecursive(root.Right)
}
