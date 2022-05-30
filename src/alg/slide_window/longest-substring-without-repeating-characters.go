package slidewindow

import "fmt"

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	result := 0
	left := 0
	right := 0
	for right <= len(s)-1 {
		// 检查条件是否合格
		v := s[right]
		if c, ok := window[v]; ok {
			window[v] = c + 1
		} else {
			window[v] = 1
		}
		right++
		for window[v] > 1 {
			window[s[left]] -= 1 // remove last
			left++
		}
		tmp := right - left
		fmt.Println(right, left)
		if tmp > result {
			result = tmp // get max value
		}
	}
	return result
}

/**
1-0
2-0
3-0
4-1
5-1
4
42
*/
