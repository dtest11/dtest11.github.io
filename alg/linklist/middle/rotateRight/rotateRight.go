package rotateRight

import "github.com/dtest11/alg/linklist/build"

type ListNode = build.ListNode

// TODO
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 0
	tail := head
	for tail != nil {
		n++
		tail = tail.Next
	}
	add := n - k%n
	if add == n {
		return head
	}
	tail.Next = head // 连接收尾形成环
	for ; add > 0; add-- {
		tail = tail.Next
	}
	result := tail.Next
	tail.Next = nil
	return result
}
