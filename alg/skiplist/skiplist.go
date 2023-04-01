package skiplist

import (
	"math/rand"
	"time"
)

const (
	MaxL = 32
	P    = 0.5
)

type Node struct {
	v  int      // value
	ls []*Level // node's index
}

type Level struct {
	next *Node
}

type SkipList struct {
	hn *Node // header
	h  int   // height
	c  int   // count
}

func New() *SkipList {
	return &SkipList{
		hn: NewNode(MaxL, 0), // head node doesn't count
		h:  1,
		c:  0,
	}
}

func NewNode(level, val int) *Node {
	node := new(Node)
	node.v = val
	node.ls = make([]*Level, level)

	for i := 0; i < len(node.ls); i++ {
		node.ls[i] = new(Level)
	}
	return node
}

/*
*
get a random level, the odds is

	1/2  = 1
	1/4  = 2
	1/8  = 3
	...
	1/2 + 1/4 + 1/8 ... + 1/n
*/
func (sl *SkipList) randomL() int {
	l := 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for r.Float64() < P && l < MaxL {
		l++
	}
	return l
}

func (sl *SkipList) insert(value int) bool {
	update := make([]*Node, MaxL)

	th := sl.hn
	for i := sl.h - 1; i >= 0; i-- {
		// 查找每一次的正确位置
		for th.ls[i].next != nil && th.ls[i].next.v < value {
			th = th.ls[i].next
		}

		// make sure no duplicates
		if th.ls[i].next != nil && th.ls[i].next.v == value {
			return false
		}

		update[i] = th
	}

	// get the level for the new item
	level := sl.randomL()
	node := NewNode(level, value)

	if level > sl.h {
		sl.h = level
	}

	for i := 0; i < level; i++ {
		// concat new node of the upper level to the head node
		if update[i] == nil {
			sl.hn.ls[i].next = node
			continue
		}
		// insert node
		node.ls[i].next = update[i].ls[i].next
		update[i].ls[i].next = node
	}
	return true
}

//func (sl *SkipList) query(value []int) bool {
//	th := sl.hn
//	// search for the top level
//	for i := sl.h - 1; i >= 0; i-- {
//		for th.ls[i].next != nil && th.ls[i].next.v <= value {
//			th = th.ls[i].next
//		}
//		// if no match, then search the next level
//		if th.v == value {
//			node = th
//			break
//		}
//	}
//
//	if node == nil {
//		return nil, false
//	}
//}
