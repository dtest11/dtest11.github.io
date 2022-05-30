package sort

import (
	"fmt"
	"testing"
)

var data = []int{9, 8, 10, 3, 6}

func TestInsertSort(t *testing.T) {
	//InsertSort(data)
	//InsertSort1(data)
	result :=MergeSort(data)
	fmt.Println(result)
}
