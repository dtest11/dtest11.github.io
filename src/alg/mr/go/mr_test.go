package mr

import (
	"fmt"
	"testing"
)

func Test_ms(t *testing.T) {
	//var nums = []int{3, 2, 4, 1}
	//nums = mergeSort(nums)
	//fmt.Println(nums)
	nums1 := []int{0, 5, 2, 3, 1, 8}
	fmt.Println(mergeSort(nums1))
}
