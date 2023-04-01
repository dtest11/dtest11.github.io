package removeDuplicateNodes

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeSortDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var dummy = &ListNode{Val: -1, Next: head}
	var tmp = dummy
	for dummy.Next != nil && dummy.Next.Next != nil {
		if dummy.Next.Val == dummy.Next.Next.Val {
			dummy.Next = dummy.Next.Next
		} else {
			dummy = dummy.Next
		}
	}
	return tmp.Next
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	exist := make(map[int]struct{})
	exist[head.Val] = struct{}{}
	tmp := head
	for tmp.Next != nil {
		cur := tmp.Next
		if _, ok := exist[cur.Val]; ok {
			tmp.Next = tmp.Next.Next
		} else {
			exist[cur.Val] = struct{}{}
			tmp = tmp.Next
		}
	}
	return head
}
