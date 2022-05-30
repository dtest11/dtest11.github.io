package array

/*****
Example 1:

Input: [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Example 2:

Input: [-7,-3,2,3,11]
Output: [4,9,9,49,121]

***/

// 1. 对数组里面的每一一个元素 去绝对值， 然后排序
func sortedSquaresS1(A []int) []int {
	for index, val := range A {
		A[index] = abs(val)
	}
	// 来一次冒泡排序

	for i := 0; i < len(A); i++ {
		for j := 1; j < len(A); j++ {
			if A[j] < A[j-1] {
				A[j], A[j-1] = A[j-1], A[j]
			}
		}
	}
	return A
}

// 因为数组 A 已经排好序了， 所以可以说数组中的负数已经按照平方值降序排好了，数组中的非负数已经按照平方值升序排好了。
// 举一个例子，若给定数组为 [-3, -2, -1, 4, 5, 6]，数组中负数部分 [-3, -2, -1] 的平方为 [9, 4, 1]，
// 数组中非负部分 [4, 5, 6] 的平方为 [16, 25, 36]。我们的策略就是从前向后遍历数组中的非负数部分，并且反向遍历数组中的负数部分。
//  左右 2 个游标 ，
// 通過建立左右兩個指針，然後比較左右兩個指針所指向的值的絕對值那個更大，更大的那個顯然就應該放到後面，然後接着比較即可。
// 这个题目其实 是用2 个指针来排序的
func sortedSquares(A []int) []int {
	i := 0
	j := len(A) - 1
	copyA := make([]int, len(A))
	for index := len(A) - 1; index >= 0; index-- {
		if abs(A[i]) < abs(A[j]) {
			copyA[index] = A[j] * A[j]
			j = j - 1 //  现在已经取的了最大值
		} else {
			copyA[index] = A[i] * A[i]
			i = i + 1
		}
	}
	return copyA
}

func abs(input int) int {
	if input < 0 {
		return -input
	}
	return input
}

//func main() {
//	//datas := []int{1, 6, 2, 4, 3, 0}
//	datas := []int{-4, -1, 0, 3, 10}
//	a := sortedSquares(datas)
//	fmt.Println(a)
//}
