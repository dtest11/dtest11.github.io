package sort

// MergeSort 归并排序
//归并排序（MERGE-SORT）是利用归并的思想实现的排序方法，该算法采用经典的分治（divide-and-conquer）策略
// （分治法将问题分(divide)成一些小的问题然后递归求解，
//而治(conquer)的阶段则将分的阶段得到的各答案"修补"在一起，即分而治之)。
func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	//递[归]
	middle := len(data) / 2
	//不断地进行左右对半划分
	left := MergeSort(data[:middle])
	right := MergeSort(data[middle:])
	//合[并]
	return merge(left, right)
}
func merge(left, right []int) (result []int) {
	l, r := 0, 0
	// 注意：[左右]对比，是指左的第一个元素，与右边的第一个元素进行对比，哪个小，就先放到结果的第一位，然后左或右取出了元素的那边的索引进行++
	for l < len(left) && r < len(right) {
		//从小到大排序.
		if left[l] > right[r] {
			result = append(result, right[r])
			//因为处理了右边的第r个元素，所以r的指针要向前移动一个单位
			r++
		} else {
			result = append(result, left[l])
			//因为处理了左边的第r个元素，所以r的指针要向前移动一个单位
			l++
		}
	}
	// 比较完后，还要分别将左，右的剩余的元素，追加到结果列的后面(不然就漏咯）。
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}
