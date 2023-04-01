package has_circle

import "github.com/dtest11/alg/linklist/build"

type ListNode = build.ListNode

// hasCycle 快慢指针，是否相遇，注意判断停止的条件
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// detectCycle 先检查是否有环
// 在发现环的地方停止，然后一个节点从开头开始遍历，每个节点每次走一步，然后相遇的地方就是环的开始节点
func detectCycle(head *ListNode) *ListNode {
	if !hasCycle(head) {
		return nil
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	slow = head
	for slow != nil && fast != nil {
		if slow == fast {
			return slow
		}
		slow = slow.Next
		fast = fast.Next
	}
	return nil
}
