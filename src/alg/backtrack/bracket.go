package main

func generateBrackets(n int) [][]string {
	var result [][]string
	var track []string
	generateBracketBackTrack(n, &result, track, 0, n, n)
	return result
}

func generateBracketBackTrack(n int, result *[][]string, track []string, i int, left int, right int) {
	// 结束条件
	// 如果是合法的括号组合的话
	if i == 2*n {
		if right < left {
			return // 不合法
		}
		if left < 0 || right < 0 {
			return
		}
		if left == right {
			tmp := make([]string, n)
			copy(tmp, track)
			*result = append(*result, tmp)
			return
		}
	}
	var choices = []string{"(", ")"}
	for _, v := range choices {
		track = append(track, v) // 添加选择
		if v == "(" {
			generateBracketBackTrack(n, result, track, i+1, left-1, right)
		} else {
			generateBracketBackTrack(n, result, track, i+1, left, right-1)
		}

		track = track[:len(track)-1]
	}
}
