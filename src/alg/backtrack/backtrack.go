package main

import "fmt"

func permutes(nums []int) [][]int {
	var track []int
	currentIndex := 0
	numsLen := len(nums)
	var result [][]int
	permutesHelper(nums, track, currentIndex, numsLen, &result)
	return result
}

func permutesHelper(nums []int, track []int, currentIndex int, numsLen int, result *[][]int) {
	if len(track) == numsLen {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*result = append(*result, tmp)
		return
	}
	for i := currentIndex; i < numsLen; i++ {
		track = append(track, nums[i])
		permutesHelper(nums, track, currentIndex+1, numsLen, result)
		track = track[:len(track)-1]
	}
}

/***
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。



示例 1：

输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combinations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func combine(n int, k int) [][]int {
	var track []int
	var result [][]int
	combineHelper(n, k, 1, &result, track)
	return result
}

func combineHelper(n int, k int, start int,
	result *[][]int, track []int) {

	if len(track) == k {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*result = append(*result, tmp)
		fmt.Println(tmp)
		return
	}
	for i := start; i <= n; i++ {
		// 添加选择
		track = append(track, i)
		combineHelper(n, k, i+1, result, track)
		// 撤销选择
		track = track[:len(track)-1]
	}
}

func main() {
	nums := []int{1, 2, 3, 4}
	for _, v := range permutes(nums) {
		fmt.Println(v)
	}
}
