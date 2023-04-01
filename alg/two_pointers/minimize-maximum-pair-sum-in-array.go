package twopointers

import "sort"

// 请你在最优数对划分的方案下，返回最小的 最大数对和 。
/**
题解： 对于数组a<b<c<d, 那么（a,d),(b,c) 形成的最大数组和，会出现最小的，
如果: a,b/ cd ，那么 不会是最大的数组里面最小和

[3,5,2,3]=》 2,3,3,5
2,5 / 3, 3 = max(7,6)=7
2,3 / 3,5 =max(5,8) =8 是大的数组和

https://leetcode.cn/problems/minimize-maximum-pair-sum-in-array/
**/

func minPairSum(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	l := 0
	r := len(nums) - 1
	var res = 0
	for l < r {
		sum := nums[l] + nums[r]
		if sum > res {
			res = sum
		}
		l = l + 1
		r = r - 1
	}
	return res
}
