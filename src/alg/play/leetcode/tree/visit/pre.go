package visit

import (
	"fmt"
	"log"

	"github.com/9999-dollor/algorithm/define"
	"github.com/9999-dollor/algorithm/leetcode/stack"
	"github.com/9999-dollor/algorithm/leetcode/tree"
)

/***
tree 的遍历都是 root ->left -> right ,

栈S;
p= root;
while(p || S不空){
    while(p){
        访问p节点；
        p的右子树入S;
        p = p的左子树;
    }
    p = S栈顶弹出;
}
***/
func PreOrder(root *define.TreeNode) {
	if root == nil {
		return
	}
	cur := root
	var stack []*define.TreeNode
	for len(stack) != 0 || cur != nil {
		for cur != nil {
			fmt.Printf("%d ", cur.Val)
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
}

func PreorderRecursive(root *define.TreeNode) {
	if root == nil {
		return
	}

	PreorderRecursive(root.Left)
	PreorderRecursive(root.Right)
}

// PreOrderStackV1 这里我们使用了栈，来模拟循环， 栈是先进后出， 因为我们是
//  root - left -right => 因此入栈的顺序，需要先，right , left ,通过这个来保证left 先被弹出来
func PreOrderStackV1(root *define.TreeNode) {
	var stack = []*define.TreeNode{}
	if root == nil {
		return
	}
	stack = append(stack, root)
	for len(stack) != 0 {
		ele := stack[len(stack)-1]
		fmt.Println(ele.Val)
		stack = stack[:len(stack)-1] // pop element
		if ele.Right != nil {
			stack = append(stack, ele.Right)
		}
		if ele.Left != nil {
			stack = append(stack, ele.Left)
		}
	}
}

func PreOrderV2(root *tree.TreeNode) {
	if root == nil {
		return
	}
	stk := stack.NewStack()
	current := root
	for !stk.Empty() || current != nil {
		for current != nil {
			log.Printf("%v", current.Val)
			stk.Push(current)
			current = current.Left
		}

		if !stk.Empty() {
			current = stk.Peek().(*tree.TreeNode)
			stk.Pop()
			current = current.Right
		}
	}
}
