package main

/***
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Notice that the solution set must not contain duplicate triplets.
Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Example 2:

Input: nums = []
Output: []
Example 3:

Input: nums = [0]
Output: []
*/


func threeSum(nums []int) [][]int {
	return threeSumTarget(nums, 0)
}

func threeSumTarget(nums []int, target int) [][]int {
	var result [][]int
	data := sort(nums)
	for i := range data {
		if i > 0 && data[i].ele == data[i-1].ele {
			continue
		}
		current := data[i]
		left := i + 1
		right := len(data) - 1

		for left < right {
			sum := data[left].ele + data[right].ele + current.ele
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				res := []int{current.ele, data[left].ele, data[right].ele}
				result = append(result, res)
				for left < right  && data[left].ele  == data[left+1].ele{
					left++ // 去重重复的数字
				}
				for left < right  && data[right].ele  == data[right-1].ele{
					right-- // 去重重复的数字
				}
				left++
				right--
			}
		}
	}
	return result
}
