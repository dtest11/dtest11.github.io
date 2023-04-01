package remove_element

func deleteNode(head *ListNode, val int) *ListNode {
	var dummy = &ListNode{Val: -1, Next: head}
	var tmp = dummy
	for dummy.Next != nil {
		if dummy.Next.Val == val {
			dummy.Next = dummy.Next.Next
		} else {
			dummy = dummy.Next
		}
	}
	return tmp.Next
}
