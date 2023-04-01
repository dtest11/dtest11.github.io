package array

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	var result = make([][]int, numRows)
	result[0] = []int{1}

	if numRows == 1 {
		return result
	}
	result[1] = []int{1, 1}
	if numRows == 2 {
		return result
	}
	result[2] = []int{1, 2, 1}
	if numRows == 3 {
		return result
	}
	for i := 3; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		result[i][i] = 1
		for j := 1; j <= i-1; j++ {
			a := result[i-1][j-1]
			b := result[i-1][j]
			result[i][j] = a + b // 当前这层的值等于上面的2个累加
		}
	}
	return result
}
