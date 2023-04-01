package reversenodesinevenlengthgroups

import (
	"github.com/dtest11/alg/linklist/build"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     val int
 *     Next a*ListNode
 * }
 */

type ListNode = build.ListNode

// 翻转数据长度是偶数的节点
// https://leetcode.com/problems/reverse-nodes-in-even-length-groups/submissions/902322534/

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var result = &ListNode{Val: -1}
	var ans = result

	for j := 0; head != nil; j++ {
		var data []int
		for i := 0; i < j+1 && head != nil; i++ {
			data = append(data, head.Val)
			head = head.Next
		}
		if len(data)%2 == 0 {
			data = reverseArrays(data)
		}

		for _, v := range data {
			result.Next = &ListNode{Val: v}
			result = result.Next
		}
	}
	return ans.Next
}

func reverseArrays(data []int) []int {
	for j, i := 0, len(data)-1; j <= i; {
		data[j], data[i] = data[i], data[j]
		i--
		j++
	}
	return data
}

func reverseNode(node *ListNode) *ListNode {
	var cur = node
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
