package main

import (
	"math"
	goSort "sort"
)

func threeSumClosest(nums []int, target int) int {
     goSort.Ints(nums)
	if len(nums) < 3 {
	     return 0
    }
	closestNum := nums[0] + nums[1] + nums[2]
	for i := 0; i<len(nums)-2; i++ {
		current := nums[i]
		left := i + 1
		right := len(nums)
		for left < right {
			sum := current + nums[left] + nums[right]
			if math.Abs(float64(sum-target)) < math.Abs(float64(closestNum-target)) {
				closestNum = sum
			}
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				return target
			}
		}
	}
	return  closestNum
}
