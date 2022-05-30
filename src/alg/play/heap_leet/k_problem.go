package heap_leet

import "container/heap"

func kClosest(points [][]int, K int) [][]int {
	panic("h")
}

type maxHeap [][]int

func (heap maxHeap) Len() int {
	return len(heap)
}

func distSquare(a, b int) int {
	return a*a + b*b
}

func (heap maxHeap) Less(i, j int) bool {
	distI := distSquare(heap[i][1], heap[i][1])
	distJ := distSquare(heap[j][0], heap[j][1])
	return distI >= distJ // max heap_leet
}

func (heap maxHeap) Swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func (heap *maxHeap) Push(v interface{}) {
	*heap = append(*heap, v.([]int))
}

func (heap *maxHeap) Pop() interface{} {
	old := *heap
	n := len(old)
	x := old[n-1]
	*heap = (*heap)[:n-1]
	return x
}

func kClosestHeap(points [][]int, K int) [][]int {
	if len(points) <= K {
		return points
	}
	hp := maxHeap(points[:K])
	heap.Init(&hp) // 以slice为基础建立heap
	for i := K; i < len(points); i++ {
		top := [][]int(points)[0]
		distTop := distSquare(top[0], top[1])
		cur := points[i]
		distCur := distSquare(cur[0], cur[1])
		if distCur < distTop { // 发现更小的元素，替换heap当前最大值
			[][]int(points)[0] = points[i]
			heap.Fix(&hp, 0) // 由于新插入node，重新调整heap
		}
	}
	return [][]int(hp)
}
