package v2

type Element struct {
	Value any
	prev  *Element
	next  *Element
}

type Payload struct {
	key int
	val int
}

type DoublyLinkedList struct {
	head *Element
	tail *Element
	len  int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	head := &Element{Value: Payload{}}
	tail := &Element{Value: Payload{}}
	head.next = tail
	tail.prev = head
	return &DoublyLinkedList{head: head, tail: tail, len: 0}
}

func (l *DoublyLinkedList) PushFront(v interface{}) *Element {
	ele := &Element{Value: v}

	next := l.head.next

	ele.next = next
	ele.prev = l.head

	l.head.next = ele
	next.prev = ele
	l.len++
	return ele
}

func (l *DoublyLinkedList) PushBack(v interface{}) *Element {
	ele := &Element{Value: v}

	ele.next = l.tail
	ele.prev = l.tail.prev

	l.tail.prev.next = ele
	l.tail.prev = ele
	l.len++
	return ele
}

func (l *DoublyLinkedList) Remove(ele *Element) {
	if l.len == 0 {
		return
	}
	// 注意双端链表 prev,next 都需要处理
	ele.prev.next = ele.next
	ele.next.prev = ele.prev
	l.len--
}

func (l *DoublyLinkedList) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.head.next
}

func (l *DoublyLinkedList) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.tail.prev
}

func (l *DoublyLinkedList) Len() int {
	return l.len
}

// func (l *DoublyLinkedList) PrintFront() {
// 	node := l.head
// 	for node != nil {
// 		fmt.Printf("%v", node.Value)
// 		node = node.next
// 	}
// 	fmt.Println()
// }

// func (l *DoublyLinkedList) PrintTail() {
// 	node := l.tail
// 	for node != nil {
// 		fmt.Printf("%v", node.Value)
// 		node = node.prev
// 	}
// 	fmt.Println()
// }
