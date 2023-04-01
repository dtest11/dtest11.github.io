package palindrome_linklist

import "fmt"

// 回文链表
type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkList struct {
	head *ListNode
}

func (l *LinkList) InsertFront(val int) {
	if l.head == nil {
		l.head = &ListNode{Val: val}
		return
	}
	copyHead := l.head
	for copyHead.Next != nil {
		copyHead = copyHead.Next
	}
	copyHead.Next = &ListNode{Val: val}
}

func (l *LinkList) Print() {
	copyHead := l.head
	for l.head != nil {
		fmt.Printf("%d \t", l.head.Val)
		l.head = l.head.Next
	}
	fmt.Println()
	l.head = copyHead
}

// 先将前半部分的链表反转
// 然后一个一个的比较 反转和剩下来的元素是否相等
func reverse(head *ListNode) *ListNode {
	var current = head
	var prev *ListNode

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	// 找到前半部分链表的尾节点并反转后半部分链表
	firstHalfEnd := middleNode(head)
	secondHalfStart := reverse(firstHalfEnd.Next)

	// 判断是否回文
	p1 := head
	p2 := secondHalfStart
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 还原链表并返回结果
	firstHalfEnd.Next = reverse(secondHalfStart)
	return true
}

func isPalindromeV2(head *ListNode) bool {
	var tmp []int
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	if len(tmp) == 0 {
		return true
	}
	i := 0
	j := len(tmp) - 1
	for i < j && i != j {
		if tmp[i] != tmp[j] {
			return false
		}
		j--
		i++
	}
	return true
}
