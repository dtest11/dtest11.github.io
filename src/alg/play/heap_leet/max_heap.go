package heap_leet

/***
Arr[(i-1)/2] Returns the parent node.
Arr[(2*i)+1] Returns the left child node.
Arr[(2*i)+2] Returns the right child node.
*/

import "errors"

// 给了一个无序数组，可能有重复数字，找到第 k 个最大的元素并且返回这个元素值。

type heapP struct {
	Array    []int
	Len      int // 当前的数量
	Capacity int // 固定最多能容纳的数量
}

// Left 查找left 节点
func (h *heapP) Left(index int) int {
	i := 2 * index
	return i
}

// Right 查找right 节点
func (h *heapP) Right(index int) int {
	i := 2*index + 1
	return i
}

// Parent 查找Parent
func (h *heapP) Parent(index int) int {
	return index / 2
}

func (h *heapP) Add(data int) error {
	if h.Len == h.Capacity-1 {
		return errors.New("超过容量")
	}
	h.Len++
	h.Array[h.Len] = data
	h.Up(h.Len) // 调整堆的位置
	return nil
}

func (h *heapP) Up(index int) {
	for {
		val := h.Array[index]
		if index >= h.Len {
			break
		}
		pIndex := h.Parent(index)
		pVal := h.Array[pIndex]
		if pIndex == index || pVal > val {
			break
		}
		h.Array[pIndex], h.Array[index] = h.Array[index], h.Array[pIndex]
		index = pIndex
	}
}

func NewHeap(capacity int) *heapP {
	return &heapP{Capacity: capacity, Array: make([]int, capacity, capacity)}
}
