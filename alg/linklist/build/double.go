package build

import "fmt"

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}
type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

/*
**
画图分析下。 注意连接节点端前后
*/
func newLinkList() *DoublyLinkedList {
	head := &Node{Val: -99}
	tail := &Node{Val: -999}
	l := &DoublyLinkedList{head: head, tail: tail}
	l.head.Next = l.tail
	l.tail.Prev = l.head

	return l
}
func (l *DoublyLinkedList) Print() {
	node := l.head.Next
	for node.Next != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
	fmt.Println()
}
func (l *DoublyLinkedList) InsertFront(val int) {
	node := &Node{Val: val}
	node.Next = l.head.Next
	node.Prev = l.head
	l.head.Next.Prev = node // 连接2个节点，插入node节点
	l.head.Next = node
	l.size++
}

func (l *DoublyLinkedList) InsertLast(val int) {
	node := &Node{Val: val}
	node.Prev = l.tail.Prev
	node.Next = l.tail
	l.tail.Prev.Next = node // 连接2个节点，插入node节点,left,right
	l.tail.Prev = node
	l.size++
}

func (l *DoublyLinkedList) IsEmpty() bool {
	return l.size == 0
}
func (l *DoublyLinkedList) DeleteFront() {
	if l.IsEmpty() {
		return
	}
	node := l.head.Next
	l.head.Next = node.Next
	node.Next.Prev = l.head
	l.size--
}
func (l *DoublyLinkedList) DeleteLast() {
	if l.IsEmpty() {
		return
	}
	node := l.tail.Prev
	node.Prev.Next = l.tail
	l.tail.Prev = node.Prev
	l.size--
}
func (l *DoublyLinkedList) GetSize() int {
	return l.size
}
