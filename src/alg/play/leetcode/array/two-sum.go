package array

// Easy
// 滑动窗口类型

/***


Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

***/

func twoSum(nums []int, target int) []int {
	var collections = make(map[int]int)
	for index, i := range nums {
		collections[i] = index
	}
	for key, v := range collections {
		v1 := target - key
		key2, ok := collections[v1]
		if ok {
			return []int{v, key2}
		}
	}

	return nil
}
