package add_two_numbers

import (
	"github.com/dtest11/alg/linklist/build"
)

/*
*
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
*/
type ListNode = build.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	var result = dummy
	carry := 0
	for l1 != nil && l2 != nil {
		val := l2.Val + l1.Val + carry
		dummy.Next = &ListNode{Val: val % 10}
		dummy = dummy.Next
		carry = val / 10
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		val := l1.Val + carry
		dummy.Next = &ListNode{Val: val % 10}
		dummy = dummy.Next
		l1 = l1.Next
		carry = val / 10
	}
	for l2 != nil {
		val := l2.Val + carry
		dummy.Next = &ListNode{Val: val % 10}
		dummy = dummy.Next
		l2 = l2.Next
		carry = val / 10
	}
	if carry != 0 {
		dummy.Next = &ListNode{Val: carry}
		dummy = dummy.Next
	}
	return result.Next
}

func reverse(root *ListNode) *ListNode {
	cur := root
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	var pre *ListNode

	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = slow.Next
	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     val int
 *     Next *ListNode
 * }
 */
func addTwoNumbersV3(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverse(l1)
	l2 = reverse(l2)
	var carry int
	var dummy = &ListNode{Val: -1}
	var result = dummy
	for l1 != nil && l2 != nil {
		temp := l1.Val + l2.Val + carry
		carry = temp / 10
		pos := temp % 10
		dummy.Next = &ListNode{Val: pos}
		l1 = l1.Next
		l2 = l2.Next
		dummy = dummy.Next
	}
	for l1 != nil {
		temp := l1.Val + carry
		carry = temp / 10
		pos := temp % 10
		dummy.Next = &ListNode{Val: pos}
		l1 = l1.Next
		dummy = dummy.Next

	}

	for l2 != nil {
		temp := l2.Val + carry
		carry = temp / 10
		pos := temp % 10
		dummy.Next = &ListNode{Val: pos}
		l2 = l2.Next
		dummy = dummy.Next

	}
	if carry != 0 {
		dummy.Next = &ListNode{Val: carry}
	}

	return reverse(result.Next)
}

/**
https://leetcode.com/problems/linked-list-components/description/
题目中需要判断链表中节点的值是否在数组 nums 中，因此我们可以使用哈希表
�
s 存储数组 nums 中的值。

遍历链表，找到第一个在哈希表
�
s 中的节点，然后从该节点开始遍历，直到遇到不在哈希表
�
s 中的节点，这样就找到了一个组件，累加答案。然后继续遍历链表，直到遍历完整个链表

作者：lcbin
链接：https://leetcode.cn/problems/linked-list-components/solution/by-lcbin-cms9/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
**/

func numComponents(head *ListNode, nums []int) int {
	var dict = make(map[int]struct{})
	for _, v := range nums {
		dict[v] = struct{}{}
	}

	var result int
	for head != nil {
		val := head.Val
		_, ok := dict[val]
		if ok {
			result++
			// 从该节点遍历，找到不再dict 中端数据
			for head != nil {
				val := head.Val
				_, ok := dict[val]
				if !ok {
					break
				}
				head = head.Next
			}
		}
		if head == nil {
			break
		}

		head = head.Next

	}
	return result
}
