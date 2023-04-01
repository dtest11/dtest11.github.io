package deleteDuplicates

import "github.com/dtest11/alg/linklist/build"

type ListNode = build.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	var dummy = &ListNode{Val: -1, Next: head}
	var tmp = dummy
	var exist = make(map[int]int)
	for head != nil {
		if v, ok := exist[head.Val]; ok {
			exist[head.Val] = v + 1
		} else {
			exist[head.Val] = 1
		}
	}

	for dummy.Next != nil {
		cur := dummy.Next.Val
		val := exist[cur]
		if val > 1 {
			dummy.Next = dummy.Next.Next
		} else {
			dummy = dummy.Next
		}
	}
	return tmp.Next
}
