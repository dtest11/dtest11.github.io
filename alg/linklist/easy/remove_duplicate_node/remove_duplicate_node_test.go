package remove_duplicate_node

import (
	"fmt"
	"testing"
)

type LinkList struct {
	head *ListNode
}

func NewLinkList() *LinkList {
	return &LinkList{
		head: &ListNode{
			Val: -1,
		},
	}
}

func (l *LinkList) Add(val int) {
	root := l.head
	for root.Next != nil {
		root = root.Next
	}
	root.Next = &ListNode{Val: val}
}

func (l *LinkList) Print(t *testing.T) {
	root := l.head
	for root != nil {
		t.Logf("%d \t", root.Val)
		root = root.Next
	}
	t.Log("done")
}

func Test_mergeTwoLists(t *testing.T) {
	a := [][]int{
		{1, 1, 2, 4, 4},
		{1, 3, 4},
	}
	var l1 = NewLinkList()
	for _, v := range a[0] {
		l1.Add(v)
	}
	l1.Print(t)

	//var l2 = NewLinkList()
	//for _, v := range a[1] {
	//	l2.Add(v)
	//}
	//l2.Print(t)

	var result = deleteDuplicates(l1.head.Next)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
