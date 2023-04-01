package lru

import "fmt"

type LRUCache struct {
	kvTable  map[int]*Node
	link     *LinkList
	capacity int
}

type Node struct {
	Key   int
	value int
	Next  *Node
	Prev  *Node
}

type LinkList struct {
	head *Node
	tail *Node
}

func newLinkList() *LinkList {
	return &LinkList{}
}

func (l *LinkList) isEmpty() bool {
	if l.head == l.tail {
		return true
	}
	return false
}

func (l *LinkList) Display() {
	head := l.head
	for head != nil {
		fmt.Printf("%v \t", head.Key)
	}
	fmt.Println()
}
func (l *LinkList) Insert2Front(node *Node) {
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}
	node.Next = l.head
	l.head.Next.Prev = node
	l.head = node
}

func (l *LinkList) Move2Front(node *Node) {
	l.RemoveNode(node)
	l.Insert2Front(node)
}

func (l *LinkList) RemoveNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (l *LinkList) RemoveTail() {
	tail := l.tail
	l.tail = tail.Prev
	l.tail.Next = nil
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		kvTable:  make(map[int]*Node, capacity),
		link:     newLinkList(),
		capacity: capacity,
	}
}

func (l *LRUCache) Get(key int) int {
	node, ok := l.kvTable[key]
	if !ok {
		return -1
	}
	l.link.Move2Front(node)
	return node.value
}

func (l *LRUCache) Put(key int, value int) {
	node, ok := l.kvTable[key]
	if ok {
		node.value = value
		l.link.Move2Front(node)
		return
	}
	node = &Node{Key: key, value: value}
	l.link.Insert2Front(node)
	if l.capacity > len(l.kvTable) {
		l.link.RemoveTail()
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
