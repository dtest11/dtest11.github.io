package sort

func QuickSort(array []int) []int {
	return quick1(array, 0, len(array)-1)
}

func quick1(array []int, low, high int) []int {
	if low < high {
		pivot := partition(array, low, high)
		quick1(array, low, pivot-1)
		quick1(array, pivot+1, high)
	}
	return array
}

func partition(array []int, low int, high int) int {
	pivot := array[high]
	for j := low; j < high; j++ {
		if array[j] < pivot {
			array[j], array[low] = array[low], array[j]
			low++
		}
	}
	array[low], array[high] = array[high], array[low]
	return low
}
