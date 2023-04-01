package design_hashmap

// ArrayLen 数组+ 链表的实现
const ArrayLen = 512

type LinkList struct {
	head *LinkNode
}

func NewLinList() *LinkList {
	return &LinkList{head: &LinkNode{Val: -1, Key: -1}}
}

type LinkNode struct {
	Key  int
	Val  int
	Next *LinkNode
}

func (l *LinkList) Get(key int) int {
	tmp := l.head
	for tmp.Next != nil {
		if tmp.Next.Key == key {
			return tmp.Next.Val
		} else {
			tmp = tmp.Next
		}
	}
	return -1
}

func (l *LinkList) Insert(key, val int) {
	tmp := l.head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &LinkNode{Val: val, Key: key}
}

func (l *LinkList) Put(key, val int) {
	tmp := l.head
	for tmp.Next != nil {
		if tmp.Next.Key == key {
			if tmp.Next.Val == val {
				return
			}
			// replace with new value
			tmp.Next.Val = val
			return
		}
		tmp = tmp.Next
	}
	tmp.Next = &LinkNode{
		Key: key,
		Val: val,
	}
}

func (l *LinkList) Del(key int) {
	tmp := l.head
	for tmp.Next != nil {
		if tmp.Next.Key == key {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
}

type MyHashMap struct {
	array [ArrayLen]*LinkList
}

func Constructor() MyHashMap {
	return MyHashMap{[512]*LinkList{}}
}

func (m *MyHashMap) Put(key int, value int) {
	pos := hash(key)
	bucket := m.array[pos]
	if bucket == nil {
		m.array[pos] = NewLinList()
		bucket = m.array[pos]
	}
	bucket.Put(key, value)
}

func (m *MyHashMap) Get(key int) int {
	pos := hash(key)
	bucket := m.array[pos]
	if bucket == nil {
		return -1
	}
	return bucket.Get(key)
}

func (m *MyHashMap) Remove(key int) {
	pos := hash(key)
	bucket := m.array[pos]
	if bucket == nil {
		return
	}
	bucket.Del(key)
}

func hash(key int) int {
	return key % ArrayLen
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
