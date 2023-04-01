package build

import "fmt"

type LinkList struct {
	Head *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinKList() *LinkList {
	return &LinkList{Head: &ListNode{Val: -1}}
}

func (l *LinkList) InsertFront(val int) {
	tmp := l.Head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &ListNode{Val: val}
}

func (l *LinkList) InsertFronts(val []int) {
	for _, v := range val {
		l.InsertFront(v)
	}
}

func (l *LinkList) ClearWith(val []int) {
	l.Head.Next = nil
	l.InsertFronts(val)
}

func (l *LinkList) Println() {
	tmp := l.Head.Next
	for tmp.Next != nil {
		fmt.Printf("%d \t", tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println()
}
