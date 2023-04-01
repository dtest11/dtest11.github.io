package array

// 2个数组的后面从最大的值开始,从后面填充nums1
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	idx := m + n - 1
	for p1 >= 0 && p2 >= 0 {
		val1 := nums1[p1]
		val2 := nums2[p2]
		if val1 > val2 {
			nums1[idx] = val1
			p1--
		} else {
			nums1[idx] = val2
			p2--
		}
		idx--
	}

	for p1 >= 0 {
		nums1[idx] = nums1[p1]
		p1--
		idx--
	}
	for p2 >= 0 {
		nums1[idx] = nums2[p2]
		p2--
		idx--
	}
}
