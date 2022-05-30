package heap_leet

import (
	"fmt"
)

func swap(A []int, i, j int) {
	A[i], A[j] = A[j], A[i]
}

//// heapify 最大堆的性质, 建立一个最大堆
//// A[left] <A[i] && A[right] < A[i]
//通过下标进行最大值比较过程
func heapify(arr []int, heapSize, i int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < heapSize && arr[left] > arr[largest] {
		largest = left
	}
	if right < heapSize && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		heapify(arr, largest, heapSize)
	}
}

func heapSort(A []int) {
	n := len(A)
	for i := n / 2; i >= 0; i-- { // 构建最大堆
		heapify(A, n, i)
	}
	for i := n - 1; i > 0; i-- {
		// Move current root to end
		swap(A, 0, i) // 把最大的数据提取出来
		heapify(A, i, 0)
	}
	fmt.Println("sort", A)
}
