package window

// func minWindow(str string, T string) string {
// 	var (
// 		left  int
// 		right int
// 		res   string = str
// 	)
// 	needs := make(map[uint8]int) // needs all need str count
// 	window := make(map[uint8]int)

// 	match := 0 // 记录 window 中已经有多少字符符合要求了

// 	for i := 0; i < len(T); i++ {
// 		elelent := T[i]
// 		if _, ok := needs[elelent]; ok {
// 			window[elelent]++ // 加入 window
// 		}
// 		b := window[elelent] == needs[elelent]
// 		if b {
// 			match++
// 		}
// 		right++

// 		if match == len(needs) { // window 中的字符串已符合 needs 的要求了
// 			res = minLen(res, window)
// 			elelent := str[left]
// 			if _, ok := needs[elelent]; ok {
// 				window[elelent]--
// 			}
// 			if window[elelent] < needs[elelent] {
// 				match--
// 			}
// 			left++
// 		}
// 	}
// 	return res
// }

func minLen(a, b string) string {
	if len(a) > len(b) {
		return b
	}
	return a
}

func size(data map[uint8]int) int {
	var res int
	for _, v := range data {
		res += v
	}
	return res
}

// func findAnagrams(s, p string) string {
// 	var (
// 		left   int
// 		right  int
// 		length        = len(s)
// 		res    string = s
// 	)
// 	needs := make(map[uint8]int) // needs all need str count
// 	window := make(map[uint8]int)

// 	match := 0 // 记录 window 中已经有多少字符符合要求了

// 	for i := 0; i < length; i++ {
// 		element := p[i]
// 		if _, ok := needs[element]; ok {
// 			window[element]++ // 加入 window
// 		}
// 		if window[element] == needs[element] { // 该元素的数量已经满足
// 			match++
// 			jd
// 			right++ // 移动位置

// 			for match == size(needs) { // window 中的字符串已符合 needs 的要求了
// 				res = minLen(res, window)
// 				element := s[left]
// 				if _, ok := needs[element]; ok {
// 					window[element]--
// 				}
// 				if window[element] < needs[element] {
// 					match--
// 				}
// 				left++
// 			}
// 		}
// 		return res
// 	}

// }
