package dynamic_programming

import (
	"golang.org/x/exp/constraints"
)

/*
*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。
*/
func maxSubArray(nums []int) int {
	res := 0
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		//temp := dp[i-1] + nums[i]
		//if temp > nums[i] {
		//	dp[i] = temp
		//} else {
		//	dp[i] = nums[i]
		//}
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	}
	res = dp[0]
	for _, v := range dp {
		if v > res {
			res = v
		}
	}
	return res
}

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
