package construct

import (
	"fmt"
	"testing"
)

// medium

// post order:后序遍历 left-right - root
// in order: 中序遍历 left -root - right
// 通过postorder 的最后一次元素位置， 可以将inorder 切分成2部分，left ,right tree
// 然后递归的遍历 构建left,right tree 的结构，

// TODO bug 还没有完成写对
func buildTreeFromInorderPostOrder(inorder []int, postorder []int) *TreeNode {
//func buildTreeFromInorderPostOrder(inorder []int, postorder []int) *TreeNode {
	// base case
	if len(inorder) == 0 {
		return nil
	}
	// base case len =1
	if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}
	index := make(map[int]int)
	for i, v := range inorder {
		index[v] = i
	}
	//return inorderPostOrderHelper(inorder, postorder, length, index, 0, length, 0, length)
	var pIndex = len(inorder)-1
	return buildUtil(inorder,postorder,0,len(inorder)-1,index,&pIndex)
}
/**

// inorderPostOrderHelper 遵循 左开右闭的原则，  a[left:right]
func inorderPostOrderHelper(inorder []int,
	postorder []int, length int, index map[int]int,
	iL, iR, pL, pR int) *TreeNode {
	// base case
	if len(inorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}

	if pR > length || iR > length {
		return nil
	}
	if pL > pR {
		return nil
	}
	if iL > iR {
		return nil
	}

	//lastIdx := len(postorder) - 1
	val := postorder[pR]
	pR = pR - 1
	// 这里切换postorder / inorder left, right
	idxInorder := index[val]

	node := &TreeNode{Val: val}
	node.Left = inorderPostOrderHelper(inorder, postorder, length, index, iL, idxInorder-1, pL, pL+idxInorder-1)
	node.Right = inorderPostOrderHelper(inorder, postorder, length, index, idxInorder+1, iR, pR-idxInorder+1, pR)
	return node
}
 */

// passed
// https://www.geeksforgeeks.org/construct-a-binary-tree-from-postorder-and-inorder/
func buildUtil(inorder []int, postorder []int, inStart, inEnd int, index map[int]int, postIndex *int) *TreeNode {
	if inStart > inEnd {
		return nil
	}
	val := postorder[*postIndex]
	node := &TreeNode{Val: postorder[*postIndex]}
	fmt.Println(*postIndex)
	*postIndex = *postIndex - 1

	if inStart == inEnd {
		return node
	}
	idx := index[val]
	node.Right = buildUtil(inorder, postorder, idx+1, inEnd, index, postIndex)
	node.Left = buildUtil(inorder, postorder, inStart, idx-1, index, postIndex)
	return node
}

type args struct {
	inorder   []int
	postOrder []int
}

func Test_b(t *testing.T) {
	var arg = args{
		inorder:   []int{9, 3, 15, 20, 7},
		postOrder: []int{9, 15, 7, 20, 3},
	}
	result := buildTreeFromInorderPostOrder(arg.inorder, arg.postOrder)
	fmt.Println(result)
}
