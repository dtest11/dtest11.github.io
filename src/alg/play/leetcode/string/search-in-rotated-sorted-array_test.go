package string

import (
	"fmt"
	"testing"
)

/**
go test -v github.com/9999-dollor/algorithm/leetcode/string -run=Test_search
*/

func search(nums []int, target int) int {
	// base case
	if nums == nil {
		return -1
	}
	if len(nums) == 1 && target != nums[0] {
		return -1
	}
	index := 0
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		prev := nums[i-1]
		if prev > cur {
			index = i
			break
		}
	}

	if nums[index] == target {
		return index
	}
	arrayPrev := nums[0:index]
	arrayLast := nums[index:]
	l := 0
	r := len(arrayPrev) - 1
	for l <= r {
		mid := l + (r-l)/2
		if arrayPrev[mid] == target {
			return mid
		} else if arrayPrev[mid] > target {
			// 4,5,6,7
			r = mid - 1
		} else {
			// nums[mid] < target
			l = mid + 1
		}
	}

	l = 0
	r = len(arrayLast) - 1
	for l <= r {
		mid := l + (r-l)/2
		if arrayLast[mid] == target {
			return index + mid
		} else if arrayLast[mid] > target {
			r = mid - 1
		} else {
			// arrayLast[mid] < target
			l = mid + 1
		}
	}
	return -1
}

func Test_search(t *testing.T) {
	var source = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(source, 2))

	source = []int{1}
	fmt.Println(search(source, 0))

	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "search", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 6}, want: 2},
		{name: "search1", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 2}, want: 6},
		{name: "search2", args: args{nums: []int{0}, target: 1}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("%s  got %v, want %v", tt.name, got, tt.want)
			}
			//t.Logf("result:%v", tt.args.nums)
		})
	}
}

/**




There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).

Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].

Given the array nums after the rotation and an integer target, return true if target is in nums, or false if it is not in nums.

You must decrease the overall operation steps as much as possible.



Example 1:

Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true
Example 2:

Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false
**/
