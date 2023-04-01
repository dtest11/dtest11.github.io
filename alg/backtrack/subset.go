package main

// subsets 输入[1,2,3,4] 输出所有这个集合的子集
// [1],[2]....,[1,2,3],etc
// 还是使用回溯算法
func subsets(nums []int) [][]int {
	var result [][]int
	var track []int
	subsetsBackTrack(nums, &result, track, 0)
	return result
}

func subsetsBackTrack(nums []int, result *[][]int, track []int, cur int) {
	// 这里所有的数据都满足条件
	tmp := make([]int, len(track))
	copy(tmp, track)
	*result = append(*result, tmp)
	for i := cur; i < len(nums); i++ {
		track = append(track, nums[i])
		subsetsBackTrack(nums, result, track, i+1)
		track = track[:len(track)-1]
	}
}
