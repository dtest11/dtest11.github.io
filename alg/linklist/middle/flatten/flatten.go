package flatten

import (
	"github.com/dtest11/alg/linklist/build"
	"github.com/dtest11/alg/tree"
)

/**
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。

**/

type ListNode = build.ListNode

type TreeNode = tree.TreeNode

func flatten(root *TreeNode) {
	tmp := preorderTraversal(root)
	for i := 1; i < len(tmp); i++ {
		prev, cur := tmp[i-1], tmp[i]
		prev.Left = nil
		prev.Right = cur
	}
}

func preorderTraversal(root *TreeNode) []*TreeNode {
	var result []*TreeNode
	if root == nil {
		return result
	}
	result = append(result, root)
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)
	return result
}

/**
https://leetcode.cn/problems/flatten-a-multilevel-doubly-linked-list/

**/

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

// flattenChild 将child 节点合并到Next节点，形成只有next,prev的链表
func flattenChild(root *Node) *Node {
	cur := root
	var result = cur
	for cur != nil {
		next := cur.Next
		if cur.Child != nil { //如果有孩子节点，处理孩子节点
			child := flattenChild(cur.Child)
			child.Prev = cur // 插入child节点在cur 之后
			cur.Next = child
			// 找到child的最后一个节点，
			for child.Next != nil {
				child = child.Next
			}
			if next != nil { //将child的最后一个节点插入在cur.Next 之前
				child.Next = next
				next.Prev = child
			}
			cur.Child = nil // 将child 置空
		}

		cur = next // 调整循环的位置，从cur.Next节点遍历
	}
	return result
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	idx := 1
	var h = head
	var m = &ListNode{Val: -1}
	var mRoot = m
	var l = &ListNode{Val: -1} // 存储前半部分
	var r = &ListNode{Val: -1} // 存储后半部分
	var result = l
	for head != nil {
		if idx >= left && idx <= right {
			m.Next = head // 满足条件的数据
			m = m.Next
		} else if idx < left {
			l.Next = head //前半部分数据
			l = l.Next
		} else {
			r.Next = head // 后半部分
			break
		}
		idx++
		head = head.Next
	}
	l.Next = nil
	m.Next = nil
	if mRoot.Next == nil {
		return h
	}
	mRoot = reverse(mRoot.Next)
	l.Next = mRoot
	for mRoot.Next != nil {
		mRoot = mRoot.Next
	}
	mRoot.Next = r.Next
	return result.Next
}

func reverse(node *ListNode) *ListNode {
	cur := node
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func rotateRight(head *ListNode, k int) *ListNode {
	l := 0
	cur := head
	data := []int{}
	for cur != nil {
		data = append(data, cur.Val)
		l++
		cur = cur.Next
	}
	k = k % l
	// 将head[i],head[i+k]的元素互换
	for i := 0; i < len(data); i++ {
		data[i], data[(i+k)%l] = data[(i+k)%l], data[i]
	}
	var dummy = &ListNode{Val: 0}
	var result = dummy
	for _, v := range data {
		dummy.Next = &ListNode{Val: v}
		dummy = dummy.Next
	}
	return result.Next

}
