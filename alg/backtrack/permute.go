package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

/**
输入一个不包含重复数字的数组 nums，返回这些数字的全部排列。

vector<vector<int>> permute(vector<int>& nums);
比如说输入数组 [1,2,3]，输出结果应该如下，顺序无所谓，不能有重复：

[
 [1,2,3],
 [1,3,2],
 [2,1,3],
 [2,3,1],
 [3,1,2],
 [3,2,1]
]
**/

func permute[T constraints.Integer](nums []T) [][]T {
	var result [][]T
	var track []T
	permuteBackTrack(nums, &result, track, 0)
	return result
}

func permuteBackTrack[T constraints.Integer](nums []T, result *[][]T, track []T, cur int) {
	// 满足条件 返回
	//fmt.Println(cur, track, len(track))
	if len(track) == len(nums) {
		tmp := make([]T, len(track))
		copy(tmp, track)
		fmt.Println(track)
		*result = append(*result, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		// 检查重复
		skip := false
		for _, v := range track {
			if v == nums[i] {
				skip = true
				break
			}
		}
		if skip {
			continue
		}
		// 添加选择
		track = append(track, nums[i])
		permuteBackTrack(nums, result, track, i+1)
		// 撤销选择
		track = track[:len(track)-1]
	}
}
