package binary

func binarySeach(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		// mid := (l + r) / 2
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return target
		}
	}
	return -1
}
