package construct

import (
	"fmt"
	"testing"

	"github.com/9999-dollor/algorithm/leetcode/tree"
)

/**
Construct Binary Tree from Preorder and Inorder Travers

1. 前序遍历 root -left -right  2 1 3
2. 中序遍历  left - root -right 1 2 3
    2

	1 3

	解法： 先从中序遍历的结果中第一个元素 找到root ， 然后根据该root 就可以将元素分成2部分(left,right)
	然后递推的 利用上面的规律 将所有的元素就最后彻底分开，

	trick: 这里利用空间换时间,map 中切分left, right
**/

type TreeNode = tree.TreeNode

// 执行超时，逻辑没有问题
func buildTreeLimitExceeded(p []int, i []int) *TreeNode {
	if len(p) == 0 && len(i) == 0 {
		return nil
	}
	root := &TreeNode{Val: p[0]}
	if len(p) > len(i) {
		return root
	}
	if len(p) == 1 && len(i) == 1 {
		fmt.Println("last:", p, i)
		return root
	}
	var idx int
	for j, v := range i { // 这里需要优化
		if v == p[0] {
			idx = j
			break
		}
	}
	fmt.Println("idx:", idx)

	leftP := p[1 : idx+1]
	rightP := p[idx+1:]
	leftI := i[:idx]
	rightI := i[idx+1:]

	fmt.Println("Left:", leftP, leftI)
	fmt.Println("Right:", rightP, rightI)
	root.Left = buildTreeLimitExceeded(leftP, leftI)
	root.Right = buildTreeLimitExceeded(rightP, rightI)
	return root
}

// buildTreeMapHelp 这个函数，每次传入全部的数据，但是记录左右的 索引位置
func buildTreeMapHelp(p []int, i []int, index map[int]int, pL, pR, iL, iR int, length int) *TreeNode {
	fmt.Println(pL, pR, iL, iR)
	if pL > length || pR > length || iL > length || iR > length {
		return nil
	}
	if pL > pR || iL > iR {
		return nil
	}
	eleP := p[pL : pR+1]
	eleI := i[iL : iR+1]
	if len(eleI) == 0 || len(eleP) == 0 {
		return nil
	}
	if len(eleP) == 0 {
		return &TreeNode{Val: eleP[0]}
	}

	val := p[pL]
	idx := index[val]
	root := &TreeNode{Val: val}

	step := idx - iL
	// 以下都是不算Val 这个元素， 左右的边界值都要计算，不能+1 ，或者-1
	// pOrder

	leftStartP := pL + 1  // porder 左边的开始
	leftEndP := pL + step // porder  左边的开始部分

	rightStartP := leftEndP + 1 // porder 右边的开始
	rightEndP := pR             // porder 右边的结束

	// inorder
	leftStartI := iL    // InOrder 左边的开始
	leftEndI := idx - 1 // InOrder 左边的开始部分

	rightStartI := idx + 1 //InOrder 右边的开始
	rightEndI := iR        // InOrder 右边的结束

	fmt.Println(rightStartP, rightEndP, rightStartI, rightEndI)
	fmt.Println(leftEndP, leftEndP, leftStartP, leftStartI, leftEndI)
	// fmt.Println("Left:", p[leftStartP:leftEndP], i[leftStartI:leftEndI])
	// fmt.Println("Right:", p[rightStartP:rightEndP], i[rightStartI:rightEndI])

	root.Left = buildTreeMapHelp(p, i, index, leftStartP, leftEndP, leftStartI, leftEndI, length)
	root.Right = buildTreeMapHelp(p, i, index, rightStartP, rightEndP, rightStartI, rightEndI, length)

	return root
}

func buildTree(p []int, i []int) *TreeNode {
	// base case
	if len(p) == 0 && len(i) == 0 {
		return nil
	}

	if len(p) == 1 && len(i) == 1 {
		return &TreeNode{Val: p[0]}
	}

	var index = make(map[int]int)
	for idx, v := range i {
		index[v] = idx
	}
	l := len(p) - 1
	return buildTreeMapHelp(p, i, index, 0, l, 0, l, l)
}

//func buildTree(preorder []int, inorder []int) *TreeNode {
//	if len(preorder) == 0 {
//		return nil
//	}
//
//	//前序遍历的第一个节点即根节点
//	res := &TreeNode{
//		Val: preorder[0],
//	}
//
//	if len(preorder) == 1 {
//		return res
//	}
//
//	// 在中序遍历中，根节点前面的即根节点的左子树，后面的即右子树
//
//	//匿名函数
//	idx := func(val int, nums []int) int {
//		for i, v := range nums {
//			if v == val {
//				return i
//			}
//		}
//		return -1
//	}(res.Val, inorder)
//
//	//递归,构建根节点左子树和右子树
//	res.Left = buildTree(preorder[1:idx+1], inorder[:idx])
//	res.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
//	return res
//}

func Test_buildTree1(t *testing.T) {
	var p = []int{3, 9, 1, 2, 20, 15, 7}
	var i = []int{1, 9, 2, 3, 15, 20, 7}
	////fmt.Println(len(p), p[:6])
	////result := buildTreeLimitExceeded(p, i)
	////fmt.Println(result)
	//
	result2 := buildTree(p, i)
	fmt.Println(result2)
	//
	p = []int{1, 2}
	i = []int{2, 1}
	////result := buildTreeLimitExceeded(p, i)
	////fmt.Println(result)
	////
	result2 = buildTree(p, i)
	fmt.Println(result2)

	p1 := []int{1, 2}
	i1 := []int{1, 2}
	result21 := buildTree(p1, i1)
	fmt.Println(result21)
}
