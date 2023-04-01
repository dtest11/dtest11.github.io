package removenodesfromlinkedlist

import (
	"github.com/dtest11/alg/linklist/build"
)

type ListNode = build.ListNode

/**
我们可以从反转后的链表头开始，像 83. 删除排序链表中的重复元素 那样，删除比当前节点值小的元素。

最后再次反转链表，即为答案。
You are given the head of a linked list.

Remove every node which has a node with a strictly greater value anywhere to the right side of it.

Return the head of the modified linked list.



Example 1:


Input: head = [5,2,13,3,8]
Output: [13,8]
Explanation: The nodes that should be removed are 5, 2 and 3.
- Node 13 is to the right of node 5.
- Node 13 is to the right of node 2.
- Node 8 is to the right of node 3.
Example 2:

Input: head = [1,1,1,1]
Output: [1,1,1,1]
Explanation: Every node has value 1, so no nodes are removed.
**/

func removeNodes(head *ListNode) *ListNode {
	dummy := reverse(head)
	var cur = dummy
	for dummy.Next != nil {
		if dummy.Val > dummy.Next.Val {
			dummy.Next = dummy.Next.Next
		} else {
			dummy = dummy.Next
		}
	}
	return reverse(cur)
}

func reverse(head *ListNode) *ListNode {
	var cur = head
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

/*
**
removeZeroSumSublists:删除连续累计和是0的元素
前缀和，对于前缀和p[i]=5,p[i+2]=5
说明i,i+2 之间的数据累加和是0，这部分的数据可以直接删除，可以把指针p[i].Next=p[i+2]
定义每个结点的前缀和是头结点到当前结点的值的和，那么我们会发现，如果两个结点p和q的前缀和相等，那么(p,q]的这些结点就是一组一组求和等于0的连续结点，就可以删除掉。
  所以我们可以用一个hash表map来存储每个结点的前缀和，如果有前缀和相等的节点自动覆盖为后面的节点，这样遍历链表时，同时求当前结点的前缀和，对每个结点cur，只要找到和它前缀和相等的结点map[sum],删除(cur,map[sum]]的操作就是cur->next = map[sum]->next。
  如果没有后面和自己前缀和相等的结点也可以，那这步就是自己的next赋值给自己的next，没有影响
*/
func removeZeroSumSublists(head *ListNode) *ListNode {
	var sumMap = make(map[int]*ListNode)
	var sum = 0
	// 因为头结点也有可能会被消掉，所以这里加一个虚拟节点作为头结点
	var dummy = &ListNode{Val: 0}
	dummy.Next = head
	var cur = dummy
	for dummy != nil {
		sum = dummy.Val + sum
		sumMap[sum] = dummy
		dummy = dummy.Next
	}
	sum = 0
	var result = cur
	for cur != nil {
		sum = sum + cur.Val
		next, ok := sumMap[sum]
		if ok {
			cur.Next = next.Next
		}
		cur = cur.Next
	}
	return result.Next
}
