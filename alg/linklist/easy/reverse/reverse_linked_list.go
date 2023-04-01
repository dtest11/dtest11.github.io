package reverse_link_list

// https://leetcode.cn/problems/reverse-linked-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表
func reversePrint(head *ListNode) []int {
	// reverse
	cur := head
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	var result []int
	// print form pre
	for prev != nil {
		result = append(result, prev.Val)
		prev = prev.Next
	}

	return result
}
