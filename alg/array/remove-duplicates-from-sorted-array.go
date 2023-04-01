package array

// removeDuplicates each unique element appears only once.
// 只能修改原来的数组，不能重新分配
func removeDuplicates(nums []int) int {
	slow := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			slow++
			nums[slow] = nums[i]
		}
	}
	return slow + 1
}
