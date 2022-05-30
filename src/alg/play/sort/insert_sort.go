package sort

import (
	"fmt"
	"log"
)

// InsertSort  . 这个我感觉更好理解和记忆
func InsertSort(array []int) {
	log.Println(array)
	for i := 0; i < len(array); i++ { // 第一张牌是有序的
		// 拿出下一张，比比较
		for j := i + 1; j < len(array) && j-1 >= 0 && array[j-1] < array[j]; {
			array[j], array[j-1] = array[j-1], array[j]
			j--
			log.Println(array)
		}
	}
	fmt.Println(array)
}

// InsertSort1 算法导论的实现
func InsertSort1(A []int) []int {
	for j := 1; j < len(A); j++ {
		key := A[j]
		// A[1, j-1]已经是有序的的数组， 将key 插入，
		i := j - 1

		// for 循环主要是将数组A[0,j-1] 比较与key 的大小，如果发现大的，
		// 那么将这个数字想右边移动以为， 空出来的就可以插入这个key
		for i > 0 && A[i] > key {
			A[i+1] = A[i]
			i = i - 1
		}
		A[j] = key

	}
	fmt.Println(A)
	return A
}
