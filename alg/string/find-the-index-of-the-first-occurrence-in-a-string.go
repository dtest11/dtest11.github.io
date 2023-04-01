package string

func strStr(haystack string, needle string) int {
	// quick check
	if len(needle) > len(haystack) {
		return -1
	}
	l := 0
	for l < len(haystack) {

		// 注意： 循环结束，下一次，需要从l+1,开始，不能从match 开始
		if haystack[l] == needle[0] {
			// 继续推进l, 和needle 查找是否匹配
			i := 0
			match := l
			for match < len(haystack) && i < len(needle) {
				if haystack[match] == needle[i] {
					match++
					i++
				} else {
					break
				}
			}
			if i == len(needle) { // 说明完全匹配
				return l
			}
		}
		l++
	}
	return -1
}
