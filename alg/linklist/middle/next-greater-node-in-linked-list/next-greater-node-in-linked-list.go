package next_greater_node_in_linked_list

import (
	"github.com/dtest11/alg/linklist/build"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     val int
 *     Next *ListNode
 * }
 */

type ListNode = build.ListNode

func nextLargerNodes(head *ListNode) []int {
	var result []int
	left := head
	for left != nil {
		curVal := left.Val
		right := left.Next
		left = left.Next
		found := false
		for right != nil {
			if right.Val > curVal {
				found = true
				result = append(result, right.Val)
				break
			}
			right = right.Next
		}
		if !found {
			result = append(result, 0)
		}
	}
	return result
}
