package mycirculardeque

import "fmt"

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}
type LinkList struct {
	head *Node
	tail *Node
	size int
}

/*
**
画图分析下。 注意连接节点端前后
*/
func newLinkList() *LinkList {
	head := &Node{Val: -99}
	tail := &Node{Val: -999}
	l := &LinkList{head: head, tail: tail}
	l.head.Next = l.tail
	l.tail.Prev = l.head

	return l
}
func (l *LinkList) Print() {
	node := l.head.Next
	for node.Next != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
	fmt.Println()
}
func (l *LinkList) InsertFront(val int) {
	node := &Node{Val: val}
	node.Next = l.head.Next
	node.Prev = l.head
	l.head.Next.Prev = node
	l.head.Next = node
	l.size++
}

func (l *LinkList) InsertLast(val int) {
	node := &Node{Val: val}
	node.Prev = l.tail.Prev
	node.Next = l.tail
	l.tail.Prev.Next = node
	l.tail.Prev = node
	l.size++
}

func (l *LinkList) IsEmpty() bool {
	return l.size == 0
}
func (l *LinkList) DeleteFront() {
	if l.IsEmpty() {
		return
	}
	node := l.head.Next
	l.head.Next = node.Next
	node.Next.Prev = l.head
	l.size--
}
func (l *LinkList) DeleteLast() {
	if l.IsEmpty() {
		return
	}
	node := l.tail.Prev
	node.Prev.Next = l.tail
	l.tail.Prev = node.Prev
	l.size--
}
func (l *LinkList) GetSize() int {
	return l.size
}

type MyCircularDeque struct {
	size int
	list *LinkList
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{size: k, list: newLinkList()}
}

func (d *MyCircularDeque) InsertFront(value int) bool {
	if d.list.GetSize() == d.size {
		return false
	}
	d.list.InsertFront(value)
	return true
}

func (d *MyCircularDeque) InsertLast(value int) bool {
	if d.list.GetSize() == d.size {
		return false
	}
	d.list.InsertLast(value)
	return true
}

func (d *MyCircularDeque) DeleteFront() bool {
	if d.list.GetSize() == 0 {
		return false
	}
	d.list.DeleteFront()
	return true
}

func (d *MyCircularDeque) DeleteLast() bool {
	if d.list.GetSize() == 0 {
		return false
	}
	d.list.DeleteLast()
	return true
}

func (d *MyCircularDeque) GetFront() int {
	if d.list.GetSize() == 0 {
		return -1
	}
	return d.list.head.Next.Val
}

func (d *MyCircularDeque) GetRear() int {
	if d.list.GetSize() == 0 {
		return -1
	}
	return d.list.tail.Prev.Val
}

func (d *MyCircularDeque) IsEmpty() bool {
	return d.list.IsEmpty()
}

func (d *MyCircularDeque) IsFull() bool {
	return d.size == d.list.GetSize()
}
