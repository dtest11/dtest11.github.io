package mr

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	mid := len(items) / 2
	left := mergeSort(items[:mid])
	right := mergeSort(items[mid:])
	return merge(left, right)
}

func merge(a []int, b []int) []int {
	final := []int{}
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
