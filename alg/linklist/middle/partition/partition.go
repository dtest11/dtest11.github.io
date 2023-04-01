package partition

import "github.com/dtest11/alg/linklist/build"

type ListNode = build.ListNode

func partition(head *ListNode, x int) *ListNode {
	// 2个node ，遍历的过程将大于x 的值，放到一个上面，小于的放到2
	// 最后合并2个链表
	left := &ListNode{Val: -1}
	right := &ListNode{Val: -1}
	rightNode := right
	leftNode := left
	for head != nil {
		if head.Val < x {
			left.Next = head
			left = left.Next
		} else {
			right.Next = head
			right = right.Next
		}
		head = head.Next
	}
	left.Next = rightNode.Next
	right.Next = nil

	return leftNode.Next
}
