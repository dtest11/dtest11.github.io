package string

import "fmt"

/**
1. 每一行不能重复
2. 每一列不能重复
3. 每个3*3 的 子数赌内没有重复的数字
**/

func isValidSudoku(board [][]byte) bool {
	record := map[string]struct{}{}
	var found = func(ele string) bool {
		if _, ok := record[ele]; ok {
			return true
		}
		record[ele] = struct{}{}
		return false
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			num := board[r][c]
			if num != '.' {
				key := fmt.Sprintf("%v+%d+r", num, r)
				key2 := fmt.Sprintf("%v+%d+c", num, c)
				key3 := fmt.Sprintf("%v+%d+%d+b", num, r/3, c/3)
				if found(key3) || found(key2) || found(key) {
					return false
				}
			}
		}
	}
	return true
}

//https://leetcode.wang/leetCode-36-Valid-Sudoku.html
