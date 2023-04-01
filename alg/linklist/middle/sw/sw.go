package sw

import (
	"github.com/dtest11/alg/linklist/build"
)

type ListNode = build.ListNode

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	m := mid(head)
	m = reverse(m)
	var result = head
	for head != nil && m != nil {
		next := head.Next
		mNext := m.Next
		m = mNext

		m.Next = nil
		head.Next = mNext
		head = next
	}
	head = result

}

func mid(root *ListNode) *ListNode {
	fast := root
	slow := root
	for ; fast != nil && fast.Next != nil; fast = fast.Next.Next {
		slow = slow.Next
	}
	return slow.Next
}

func reverse(root *ListNode) *ListNode {
	cur := root
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
