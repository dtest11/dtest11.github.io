package quick

import (
	"golang.org/x/exp/constraints"
)

func quick[v constraints.Ordered](nums []v, l int, r int) {
	if l < r {
		pivot := partition(nums, l, r)
		quick(nums, l, pivot-1)
		quick(nums, pivot+1, r)
	}
}

func partition[v constraints.Ordered](nums []v, l, r int) int {
	pivot := nums[l]
	copyL := l
	for l < r {
		for r > l && nums[r] > pivot {
			r--
		}
		for l < r && nums[l] <= pivot {
			l++
		}

		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[l], nums[copyL] = nums[copyL], nums[l]
	return l
}
