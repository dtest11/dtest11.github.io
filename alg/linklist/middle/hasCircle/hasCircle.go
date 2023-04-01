package hasCircle

import "github.com/dtest11/alg/linklist/build"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     val int
 *     Next *ListNode
 * }
 */

type ListNode = build.ListNode

func detectCycle(head *ListNode) *ListNode {
	// 检查是否有环
	slow := head
	fast := head
	var hasCircle bool
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			hasCircle = true
			break
		}
	}
	if !hasCircle {
		return nil
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
