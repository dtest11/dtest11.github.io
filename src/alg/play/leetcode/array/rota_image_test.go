package array

import "testing"

/**
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
**/
/**
1. 先从对角线旋转一次，
2. 再从水平先上反转一次， 上面，下面互换
**/

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-i][n-j] = matrix[n-i][n-j], matrix[i][j]
		}
	}
}

func Test_rorate(t *testing.T) {
	rotate(nil)
}
