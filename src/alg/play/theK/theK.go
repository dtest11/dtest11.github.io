package thek

import (
	"fmt"
	"log"
)

// FindKthLargest 从没有顺序的 的数据中，找出第K 大的值
// 快速排序的思路是这样的，在数组中随机选取一个数据，例如选取最后一个元素 m 做为分区元素，
// 比 m 小的放 m 的左边，反之放右边，再分别对左右边的分区再分别进行分区，直到分区元素
func FindKthLargest(data []int, k int) int {
	idx := partition(data, 0, len(data)-1)
	right :=len(data)-1
	left :=0
	for idx != len(data)-k {
		if idx > len(data)-k {
			idx = partition(data, left, idx-1)
		}else {
			left=idx+1
			idx=partition(data,left,right)
		}
	}
	fmt.Println(data[k])
	return data[k]

}

func QuickSort(data []int) []int {
	quickSort(data, 0, len(data)-1)
	return data
}

func partition(data []int, left, right int) int {
	x := data[right] // 最后一个元素作为比较的元素  i < x < j , 最后是这个样子的
	i := left - 1
	for j := left; j <= right-1; j++ {
		if data[j] <= x {
			i++
			data[i], data[j] = data[j], data[i] // 交换
		}
	}
	// 最后交换， data[i+1],data[right]
	data[i+1], data[right] = data[right], data[i+1]
	//log.Println(x,data)
	log.Println("x=", x, "left=", left, "right=", right, "data=", data[left:right+1])
	return i + 1
}
func quickSort(data []int, left, right int) {
	if left >= right {
		//log.Println("已经到终点")
		return
	}
	pivotIndex := partition(data, left, right)
	log.Println(data, "00000")
	quickSort(data, left, pivotIndex-1)
	quickSort(data, pivotIndex+1, right)
}
