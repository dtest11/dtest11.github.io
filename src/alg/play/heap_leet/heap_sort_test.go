package heap_leet

import (
	"fmt"
	"testing"
)

func Test_maxHeapify(t *testing.T)  {
	data :=[]int{12, 11, 13, 5, 6, 7}
	fmt.Println(len(data))
	heapSort(data)
}