package string

import "testing"

/**
4. Median of Two Sorted Arrays
Hard

10578

1574

Add to List

Share
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).



Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
Example 3:

Input: nums1 = [0,0], nums2 = [0,0]
Output: 0.00000
Example 4:

Input: nums1 = [], nums2 = [1]
Output: 1.00000
Example 5:

Input: nums1 = [2], nums2 = []
Output: 2.00000


Constraints:

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
***/

/**
https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/xun-zhao-liang-ge-you-xu-shu-zu-de-zhong-wei-s-114/
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length%2 == 0 {
		a := length/2 - 1
		b := length / 2
		res1 := kthHelp(nums1, nums2, a+1)
		res2 := kthHelp(nums1, nums2, b+1)
		return float64(res1+res2) / 2.0
	}
	b := length / 2
	res2 := kthHelp(nums1, nums2, b+1)
	return float64(res2)
}

func kthHelp(nums1 []int, nums2 []int, k int) int {
	indexA := 0
	indexB := 0
	for {
		if indexA == len(nums1) {
			return nums2[indexB+k-1]
		}
		if indexB == len(nums2) {
			return nums1[indexA+k-1]
		}
		if k == 1 {
			return min(nums1[indexA], nums2[indexB])
		}
		mid := k / 2
		// 这里防止索引越界需要处理下
		idxA := min(mid+indexA, len(nums1)) - 1
		idxB := min(mid+indexB, len(nums2)) - 1

		pivotA := nums1[idxA]
		pivotB := nums2[idxB]
		if pivotA > pivotB {
			// 丢弃nums2[0...mid]
			k = k - (idxB - indexB + 1)
			indexB = idxB + 1
		} else {
			k = k - (idxA - indexA + 1)
			indexA = idxA + 1 //丢弃nums1[0...mid]
		}
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums  []int
		numsA []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "search", args: args{nums: []int{1, 3}, numsA: []int{2}}, want: 2.000},
		{name: "search", args: args{nums: []int{1, 2}, numsA: []int{3, 4}}, want: 2.500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums, tt.args.numsA); got != tt.want {
				t.Errorf("%s  got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
