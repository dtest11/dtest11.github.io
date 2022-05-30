package visit

import (
	"testing"

	"github.com/9999-dollor/algorithm/define"
)

func prepare() *define.TreeNode {

	/*
				10
			   /  \
			 20	   30
			/ \      \
		   40  50     60
		  /
		 70
	*/

	root := &define.TreeNode{Val: 10}
	root.Left = &define.TreeNode{Val: 20}
	root.Right = &define.TreeNode{Val: 30}
	root.Left.Left = &define.TreeNode{Val: 40}
	root.Left.Right = &define.TreeNode{Val: 50}
	root.Right.Right = &define.TreeNode{Val: 60}
	root.Left.Left.Left = &define.TreeNode{Val: 70}
	return root
}
func TestPreOrder(t *testing.T) {
	root := prepare()
	PreOrder(root)
}

// func TestInorder(t *testing.T) {
// 	root := prepare()
// 	Inorder(root)
// 	log.Println("InOrderRecursive")
// 	InOrderRecursive(root)
// }
