package remove_duplicate_node

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var result = &ListNode{Val: head.Val}
	var dummy = result
	var tailVal = head.Val
	for head != nil {
		if tailVal != head.Val {
			dummy.Next = &ListNode{Val: head.Val}
			dummy = dummy.Next
			tailVal = head.Val
		}
		head = head.Next
	}
	return result
}
