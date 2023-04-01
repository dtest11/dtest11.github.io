package design_hashset

// ArrayLen 数组+ 链表的实现
const ArrayLen = 512

type MyHashSet struct {
	array [ArrayLen]*LinkList
}

type LinkList struct {
	head *LinkNode
}

func NewLinList() *LinkList {
	return &LinkList{head: &LinkNode{Val: -1}}
}

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func (l *LinkList) Contains(val int) bool {
	tmp := l.head
	for tmp.Next != nil {
		if tmp.Next.Val == val {
			return true
		} else {
			tmp = tmp.Next
		}
	}
	return false
}

func (l *LinkList) Insert(val int) {
	tmp := l.head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &LinkNode{Val: val}
}

func (l *LinkList) Del(val int) {
	if l.head == nil {
		return
	}
	tmp := l.head
	for tmp.Next != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
}

func Constructor() MyHashSet {
	return MyHashSet{array: [512]*LinkList{}}
}

func (s *MyHashSet) Add(key int) {
	if s.Contains(key) {
		return
	}
	pos := hash(key)
	bucket := s.array[pos]
	if bucket == nil {
		s.array[pos] = NewLinList()
		bucket = s.array[pos]
	}
	bucket.Insert(key)
}

func (s *MyHashSet) Remove(key int) {
	pos := hash(key)
	if s.array[pos] == nil {
		return
	}
	bucket := s.array[pos]
	bucket.Del(key)
}

func (s *MyHashSet) Contains(key int) bool {
	pos := hash(key)
	if s.array[pos] == nil {
		return false
	}
	bucket := s.array[pos]
	return bucket.Contains(key)
}

func hash(key int) int {
	return key % ArrayLen
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
