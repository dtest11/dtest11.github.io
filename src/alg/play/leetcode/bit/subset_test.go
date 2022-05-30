package bit

import (
	"fmt"
	"testing"
)

func subsetXORSum(nums []int) int {
	split := len(nums) >> 1
	fmt.Println(split)
	//var result [][]int
	for i := 0; i < len(nums); i++ {
		cur :=nums[i]
		fmt.Println(cur)
		for j := i+1; j < len(nums); j++ {
  		      next := nums[j]
  		      fmt.Println(cur,next)
		}
	}
	return split
}

func Test_subset(t *testing.T) {
	subsetXORSum([]int{1, 3})
	subsetXORSum([]int{5, 1, 6})
	//subsetXORSum([]int{5, 1, 6, 4})
}
