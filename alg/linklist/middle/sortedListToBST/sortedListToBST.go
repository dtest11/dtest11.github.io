package sortedListToBST

import (
	"github.com/dtest11/alg/linklist/build"
	"github.com/dtest11/alg/tree"
)

/***
给定一个单链表的头节点  head ，其中的元素 按升序排序 ，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差不超过 1。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
**/

type ListNode = build.ListNode

type TreeNode = tree.TreeNode

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	// 完美二叉树left < root< right
	return buildTree(result)
}

func buildTree(ele []int) *TreeNode {
	if len(ele) == 0 {
		return nil
	}
	mid := len(ele) / 2
	root := &TreeNode{Val: ele[mid]}
	root.Left = buildTree(ele[0:mid])
	root.Right = buildTree(ele[mid+1:])
	return root
}
