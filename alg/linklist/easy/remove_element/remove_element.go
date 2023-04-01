package remove_element

/**
https://leetcode.cn/problems/remove-linked-list-elements/
**/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用新链表 解法
func removeElements(head *ListNode, val int) *ListNode {
	var dummy = &ListNode{Val: -1}
	var dummyCopy = dummy
	for head != nil {
		if head.Val != val {
			dummy.Next = &ListNode{Val: head.Val}
			dummy = dummy.Next
		}
		head = head.Next
	}

	return dummyCopy.Next
}

// 使用新链表
func removeElementsV2(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	prev := dummyNode
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return dummyNode.Next
}
