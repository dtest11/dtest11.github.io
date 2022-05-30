package linklist

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
删除直接节点
要记得这个节点如果是第一个或者事最后一个， 要特殊处理一下

Remove all elements from a linked list of integers that have value val.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://www.cs.bu.edu/teaching/c/linked-list/delete/
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	current := head
	var prev *ListNode

	for current != nil {
		if current.Val == val && prev == nil {
			head = head.Next
			current = head
			continue
		}

		if current.Val == val && prev != nil {
			prev.Next = current.Next // 移除current
			current = current.Next
			continue
		}

		if current.Val != val {
			prev = current
			current = current.Next
			continue
		}
	}
	return head
}

//func main() {
	//v5 := &ListNode{Val: 1}
	//v4 := &ListNode{Val: 1}
	//v3 := &ListNode{Val: 1}
// 	v2 := &ListNode{Val: 3}
// 	v1 := &ListNode{Val: 2,Next:v2}
// 	root := &ListNode{Val: 1, Next: v1}
// 	fmt.Println(removeElements(root, 3))
//}
