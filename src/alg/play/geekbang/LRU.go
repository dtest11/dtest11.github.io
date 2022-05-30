package geekbang

import (
	"container/list"
)

/****
LRU : 算法

如果一个数据在最近一段时间没有被访问到，
那么在将来它被访问的可能性也很小。也就是说，当限定的空间已存满数据时，应当把最久没有被访问到的数据淘汰。

****/
type lru struct {
	capaticity int
	cache      map[int]*list.Element
	list       *list.List
}
type node struct {
	key int
	val int
}

func newlru(capa int) *lru {
	return &lru{capaticity: capa, cache: make(map[int]*list.Element), list: list.New()}
}

func (l *lru) get(key int) int {
	if item, ok := l.cache[key]; ok {
		l.list.MoveToFront(item)
		return item.Value.(node).val
	}
	return -1
}

func (l *lru) set(key, val int) {
	// 如果这个key ,已经存在的话,要重新放到第一位
	if item, ok := l.cache[key]; ok {
		l.list.MoveToFront(item)
		item.Value = node{key: key, val: val}
	}

	// 不存在插入的时候 ,判断容量是否已经到达上线
	if l.list.Len() >= l.capaticity {

		// 删除最后一个元素
		delete(l.cache, l.list.Back().Value.(node).val)
		l.list.Remove(l.list.Back())
	}
	// 将元素插入到第一个位置
	l.list.PushFront(node{val: val, key: key})
	l.capaticity++
	l.cache[key] = l.list.Front()
}
