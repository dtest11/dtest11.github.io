package swapPairs

import "github.com/dtest11/alg/linklist/build"

type ListNode = build.ListNode

func swapPairs(head *ListNode) *ListNode {
	var dummy = &ListNode{Next: head}
	var result = dummy
	for dummy.Next != nil && dummy.Next.Next != nil {
		tmp := dummy.Next.Next.Val
		dummy.Next.Next.Val = dummy.Next.Val
		dummy.Next.Val = tmp
		dummy = dummy.Next.Next
	}
	return result.Next
}
