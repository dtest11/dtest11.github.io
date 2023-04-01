package reverseBetween

import "github.com/dtest11/alg/linklist/build"

type ListNode = build.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     val int
 *     Next *ListNode
 * }
 */

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}

	var data []int
	var tmp = head
	n := 0
	for tmp != nil {
		data = append(data, tmp.Val)
		tmp = tmp.Next
		n++
	}

	prev := data[:left-1] // 这里需要注意
	last := data[right:]

	split := data[left-1 : right]
	if len(split) != 0 {
		i := 0
		j := len(split) - 1
		for i < j {
			split[i], split[j] = split[j], split[i]
			j--
			i++
		}
	}

	var result = &ListNode{}
	var re = result
	for _, v := range prev {
		result.Next = &ListNode{Val: v}
		result = result.Next
	}

	for _, v := range split {
		result.Next = &ListNode{Val: v}
		result = result.Next

	}

	for _, v := range last {
		result.Next = &ListNode{Val: v}
		result = result.Next

	}
	return re.Next
}
