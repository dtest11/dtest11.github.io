package heap_leet

import (
	"testing"
)

func Test_heap_Add(t *testing.T) {
	//h := NewHeap(10)
	//h.Add(14)
	//h.Add(9)
	//h.Add(12)
	//h.Add(6)
	//h.Add(1)
	//h.Add(4)
	//h.Add(3)
	//h.Add(5)
	//for _, i := range h.Container {
	//	fmt.Println(i)
	//	//fmt.Println("left",h.Container[h.LeftNode(i)],"right=",h.Container[h.RightNode(i)])
	//	//fmt.Println(h.ParentNode(i))
	//	fmt.Println(h.RightNode(i))
	//}
}

/***
有一堆石头，每块石头的重量都是正整数。

每一回合，从中选出两块 最重的 石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：

如果 x == y，那么两块石头都会被完全粉碎；
如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。

输入：[2,7,4,1,8,1]
输出：1
解释：
先选出 7 和 8，得到 1，所以数组转换为 [2,4,1,1,1]，
再选出 2 和 4，得到 2，所以数组转换为 [2,1,1,1]，
接着是 2 和 1，得到 1，所以数组转换为 [1,1,1]，
最后选出 1 和 1，得到 0，最终数组转换为 [1]，这就是最后剩下那块石头的重量。

利用priority queue， 把石头重量都存入 pq， 每次取最大两个比较，存入差值，直到pq 只剩最后一个。
*/

func MaxHeap(stones []int) {
	// build priority queue
	for i := len(stones) / 2; i >= 0; i-- { //构建最大堆
		heapify(stones, len(stones), i)
	}
}

// func lastStoneWeight(stones []int) int {
// 	MaxHeap(stones)
// 	for i := len(stones) - 1; i > 0; i-- {
// 		// Move current root to end
// 		swap(stones, 0, i) // 把最大的数据提取出来
// 		heapify(stones, i, 0)
// 	}
// 	// 每次提取
// 	i :=0
// 	for len(stones) !=1 && i <len(stones)-1{
// 		a,b :=stones[i],stones[i+1]
// 		stones=stones[2:]

// 	}
// 	return 0

// }

// 构建最大堆
func maxHeapify(A []int, i int, heapSize int) {
	largest := i
	left := 2*i + 1
	right := left + 1
	if left < heapSize && A[left] > A[i] {
		largest = left
	}
	if right < heapSize && A[right] > A[i] {
		largest = right
	}
	if largest != i {
		A[largest], A[i] = A[i], A[largest]
		maxHeapify(A, largest, heapSize)
	}
}

func TestHeap_lastStoneWeight(t *testing.T) {
	// data := []int{
	// 2, 7, 4, 1, 8, 1,
	// }
	// lastStoneWeight(data)
}
