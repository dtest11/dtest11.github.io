package spiralmatrixiv

import (
	"fmt"
	"github.com/dtest11/alg/linklist/build"
)

type ListNode = build.ListNode

// https://leetcode.cn/problems/spiral-matrix-iv/solution/2326-luo-xuan-ju-zhen-iv-by-loyair-sf4b/

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	left := 0
	right := n - 1

	top := 0
	bottom := m - 1

	var result = make([][]int, m)

	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = -1
		}
	}

	for head != nil {
		// 从左到右
		for i := left; i <= right && head != nil; i++ {
			result[top][i] = head.Val
			fmt.Printf("从左到右=[%d][%d],", top, i)
			head = head.Next
		}
		top++
		fmt.Println()
		// 从上到下
		for i := top; i <= bottom && head != nil; i++ {
			fmt.Printf("从上到下=[%d][%d],", i, right)
			result[i][right] = head.Val
			head = head.Next
		}
		right--
		fmt.Println()
		// 从右到左
		for i := right; i >= left && head != nil; i-- {
			fmt.Printf("从右到左=[%d][%d],", bottom, i)
			result[bottom][i] = head.Val
			head = head.Next
		}
		bottom--
		fmt.Println()
		// 从下到上
		for i := bottom; i >= top && head != nil; i-- {
			fmt.Printf("从下到上=[%d][%d]", i, left)
			result[i][left] = head.Val
			head = head.Next
		}
		left++
		fmt.Println()

		fmt.Println("nxt round")
	}
	return result
}
