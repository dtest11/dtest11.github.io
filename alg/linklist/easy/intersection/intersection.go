package intersection

/**
https://leetcode.cn/problems/intersection-of-two-linked-lists/
**/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 先便利存储一个链表端元素到字典，遍历B，查看B中端元素是否在字典中
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var mapA = make(map[*ListNode]struct{})
	for headA != nil {
		mapA[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := mapA[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// 双指针方式
func getIntersectionNodeV2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA := headA
	pB := headB
	for pA != pB {
		if pA != nil {
			pA = pA.Next
		} else {
			pA = headB
		}

		if pB != nil {
			pB = pB.Next
		} else {
			pB = headA
		}

	}
	return pA
}
