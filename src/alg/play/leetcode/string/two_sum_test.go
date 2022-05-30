package string

import (
	"fmt"
	"testing"
)

/***

Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution,
 and you may not use the same element twice.

You can return the answer in any order.



Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
***/

func Test_twoSum(t *testing.T) {
	source := []int{3, 2, 4}
	fmt.Println(twoSum(source, 6))
}

func twoSum(nums []int, target int) []int {
	var record = make(map[int]int)
	for index, v := range nums {
		record[v] = index
	}
	for i, v2 := range nums {
		need := target - v2
		if val, ok := record[need]; ok && i != val {
			fmt.Println("F", i, val, record, need, v2)
			return []int{i, val}
		}
	}
	return nil
}
