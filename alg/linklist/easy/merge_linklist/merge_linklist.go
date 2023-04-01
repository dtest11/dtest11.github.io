package merge_linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	dummy := &ListNode{Val: -1}
	var result = dummy
	for list1 != nil && list2 != nil {
		l1 := list1.Val
		l2 := list2.Val
		if list1.Val < list2.Val {
			dummy.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
			dummy = dummy.Next
		} else if list1.Val > list2.Val {
			dummy.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
			dummy = dummy.Next
		} else {
			dummy.Next = &ListNode{Val: l1}
			dummy.Next.Next = &ListNode{Val: l2}
			list1 = list1.Next
			list2 = list2.Next
			dummy = dummy.Next.Next
		}
	}

	if list1 != nil {
		dummy.Next = list1
	}
	if list2 != nil {
		dummy.Next = list2
	}
	return result.Next
}
