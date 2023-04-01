package removeNthFromEnd

/**
slow fast 指针解法
slow 先走n步(剩下k-n)
此时
slow 再从头部开始走
fast 从n 节点开始，当n 走到末尾(k-n)，slow 应该走了(n)
k = k-n+n
N(n-1).Next=N(n).Next

Node(n-1) -> Node(n) -> Node(n+1)
Node(n-1).Next=Node(n-1)
**/

type ListNode struct {
	Next *ListNode
	Val  int
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	fast := head
	slow := dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
