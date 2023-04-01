package sort

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func Bubble[T constraints.Ordered](data []T) {
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data)-i; j++ {
			compare := data[j] < data[j-1]
			if compare {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
		fmt.Println("sort round:", data)
	}
}

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

func mergeSort[T constraints.Ordered](items []T) []T {
	if len(items) < 2 {
		return items
	}
	mid := len(items) / 2
	left := mergeSort(items[:mid])
	right := mergeSort(items[mid:])
	return merge(left, right)
}

func merge[T constraints.Ordered](a []T, b []T) []T {
	var final []T
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
