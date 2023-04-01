package dfs

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func getTargetCopy(original *TreeNode, cloned *TreeNode, target *TreeNode, result *TreeNode) {
	if original == nil {
		return
	}
	// dfs
	// 结束条件
	if original.Val == target.Val {
		result = original
		return
	}
	// dfs 深度遍历
	getTargetCopy(original.Left, cloned, target, result)
	getTargetCopy(original.Right, cloned, target, result)
}

// 如果一个节点的父亲只有它一个儿子,则称其是”孤独节点
// 如果一个节点的父亲只有它一个儿子，则称其是”孤独节点“。要求返回所有孤独节点的值。

func getLonelyNodes(root *TreeNode, result []int) {
	if root == nil {
		return
	}
	// dfs
	// 检查是否是孤独节点
	if root.Left != nil && root.Right == nil {
		result = append(result, root.Val)
	}
	if root.Left == nil && root.Right != nil {
		result = append(result, root.Val)
	}
	getLonelyNodes(root.Left, result)
	getLonelyNodes(root.Right, result)
}

// https://leetcode.com/problems/increasing-order-search-tree/
// 二叉搜素树，Left<val<Right
// 结果只有右节点，左边都删除，并且升序
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var result TreeNode
	increasingBSTHelper(root, &result)
	return &result
}

// 中序遍历 同时构建结果二叉树
// Left <root<val
func increasingBSTHelper(root *TreeNode, result *TreeNode) {
	if root == nil {
		return
	}
	increasingBSTHelper(root.Left, result)
	root.Left = nil
	result = root
	result = root.Right
	increasingBSTHelper(root.Right, result)
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	var result = new(TreeNode)
	if root1 == nil && root2 != nil {
		result.Val = root2.Val
		result.Left = mergeTrees(nil, root2.Left)
		result.Right = mergeTrees(nil, root2.Right)
	} else if root1 != nil && root2 == nil {
		result.Val = root2.Val
		result.Left = mergeTrees(root1.Left, nil)
		result.Right = mergeTrees(root1.Right, nil)
	} else {
		result.Val = root2.Val + root1.Val
		root1.Left = mergeTrees(root1.Left, root2.Left)
		root1.Right = mergeTrees(root1.Right, root2.Right)
	}
	return result
}

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var result []int
	postorderDfs(root, &result)
	return result
}

func postorderDfs(root *Node, result *[]int) {
	if root == nil {
		return
	}
	for _, v := range root.Children {
		postorderDfs(v, result)
	}
	*result = append(*result, root.Val)
}

// https://leetcode.com/problems/n-ary-tree-preorder-traversal/
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var result []int
	preorderDfs(root, &result)
	return result
}

func preorderDfs(root *Node, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	for _, v := range root.Children {
		postorderDfs(v, result)
	}
}

// https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/
func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return sumRootToLeafDfs(root, "")
}
func sumRootToLeafDfs(root *TreeNode, path string) int {
	if root == nil {
		return 0
	}
	// leaf
	if root != nil && root.Left == nil && root.Right == nil {
		tmp := path + fmt.Sprint(root.Val)
		println(tmp)
		val, _ := strconv.ParseInt(tmp, 2, 64)
		return int(val)
	}
	left := sumRootToLeafDfs(root.Left, path+fmt.Sprint(root.Val))
	right := sumRootToLeafDfs(root.Right, path+fmt.Sprint(root.Val))
	return left + right
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left) + 1
	right := maxDepth(root.Right) + 1
	if left > right {
		return left
	}
	return right
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func maxDepthNayTree(root *Node) int {
	if root == nil {
		return 0
	}
	cur := 0
	for _, v := range root.Children {
		tmp := maxDepthNayTree(v)
		if tmp > cur {
			cur = tmp
		}
	}
	return cur + 1
}
