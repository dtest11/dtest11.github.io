package v1

import (
	"container/list"
)

type LRUCache struct {
	keyTable map[int]*list.Element
	list     *list.List
	capacity int
}

type KV struct {
	k int
	v int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		keyTable: make(map[int]*list.Element, capacity),
		list:     list.New(),
		capacity: capacity,
	}
}

func (l *LRUCache) Get(key int) int {
	ele, ok := l.keyTable[key]
	if !ok {
		return -1
	}
	l.list.MoveToFront(ele)
	return ele.Value.(*KV).v
}

func (l *LRUCache) Put(key int, value int) {
	// 检查元素是否存在
	ele, ok := l.keyTable[key]
	if ok {
		ele.Value.(*KV).v = value
		l.list.MoveToFront(ele)
		return
	}
	// 元素不存在
	ele = l.list.PushFront(&KV{k: key, v: value})
	l.keyTable[key] = ele

	if l.list.Len() > l.capacity {
		last := l.list.Back()
		delete(l.keyTable, last.Value.(*KV).k)
		l.list.Remove(last)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,val);
 */
