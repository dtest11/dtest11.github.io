package mergeInBetween

import "github.com/dtest11/alg/linklist/build"

type ListNode = build.ListNode

/*
https://leetcode.com/problems/merge-in-between-linked-lists/

You are given two linked lists: list1 and list2 of sizes n and m respectively.

Remove list1's nodes from the ath node to the bth node, and put list2 in their place.

The blue edges and nodes in the following figure indicate the result:

Build the result list and return its head.

Example 1:

Input: list1 = [0,1,2,3,4,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
Output: [0,1,2,1000000,1000001,1000002,5]
Explanation: We remove the nodes 3 and 4 and put the entire list2 in their place. The blue edges and nodes in the above figure indicate the result.
Example 2:

Input: list1 = [0,1,2,3,4,5,6], a = 2, b = 5, list2 = [1000000,1000001,1000002,1000003,1000004]
Output: [0,1,1000000,1000001,1000002,1000003,1000004,6]
Explanation: The blue edges and nodes in the above figure indicate the result.
*/
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	aNode := list1
	bNode := list1
	for i := 0; i < a-1 && aNode != nil; i++ {
		aNode = aNode.Next
	}

	for i := 0; i < b+1 && bNode != nil; i++ {
		bNode = bNode.Next
	}
	aNode.Next = list2
	tmp := list2
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = bNode
	return list1
}
