package swappingnodesinalinkedlist

import (
	"github.com/dtest11/alg/linklist/build"
)

type ListNode = build.ListNode

/*
**
You are given the head of a linked list, and an integer k.

Return the head of the linked list after swapping the values of the kth node from the beginning and the kth node from the end (the list is 1-indexed).

Example 1:

Input: head = [1,2,3,4,5], k = 2
Output: [1,4,3,2,5]
Example 2:

Input: head = [7,9,6,6,7,8,3,0,9,5], k = 5
Output: [7,9,6,6,8,7,3,0,9,5]
*/
func SwapNodes(head *ListNode, k int) *ListNode {
	fast := head
	slow := head
	result := head
	for i := 1; i < k; i++ {
		fast = fast.Next
	}
	copyFast := fast
	for ; fast.Next != nil; fast = fast.Next {
		slow = slow.Next
	}
	tmp := slow.Val
	slow.Val = copyFast.Val
	copyFast.Val = tmp
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/solution/114-er-cha-shu-zhan-kai-wei-lian-biao-by-ming-zhi-/
func flatten(root *TreeNode) {
	flatten(root.Left)
	flatten(root.Right)
	temp := root.Right
	// /把树的右边换成左边的链表
	root.Right = root.Left
	//记得要将左边置空
	root.Left = nil
	//找到树的最右边的节点
	for root.Right != nil {
		root = root.Right
	}
	//把右边的链表接到刚才树的最右边的节点
	root.Right = temp
}
