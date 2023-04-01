package getKthFromEnd

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var fast = head
	var slow = head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}

	return slow.Next

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy = &ListNode{Val: -1}
	var tmp = dummy
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			dummy.Next = &ListNode{Val: l2.Val}
			dummy = dummy.Next
			l2 = l2.Next
		} else if l1.Val < l2.Val {
			dummy.Next = &ListNode{Val: l1.Val}
			dummy = dummy.Next
			l1 = l1.Next
		} else {
			dummy.Next = &ListNode{Val: l2.Val}
			dummy = dummy.Next
			l2 = l2.Next

			dummy.Next = &ListNode{Val: l1.Val}
			dummy = dummy.Next
			l1 = l1.Next
		}
	}
	if l1 != nil {
		dummy.Next = l1
	}
	if l2 != nil {
		dummy.Next = l2
	}
	return tmp.Next
}
