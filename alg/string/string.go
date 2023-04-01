package string

/**
按照题目的描述，可以总结如下规则：

    罗马数字由 I,V,X,L,C,D,M 构成；
    当小值在大值的左边，则减小值，如 IV=5-1=4；
    当小值在大值的右边，则加小值，如 VI=5+1=6；
    由上可知，右值永远为正，因此最后一位必然为正

一言蔽之，把一个小值放在大值的左边，就是做减法，否则为加法
在代码实现上，可以往后看多一位，对比当前位与后一位的大小关系，从而确定当前位是加还是减法。当没有下一位时，做加法即可。
**/

func romanToInt(s string) int {
	sysmbols := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := 0
	for i := 0; i < len(s); i++ { // 当前数和下一个数据作比较
		// 小的数字在大的右边，直接加
		if i+1 > len(s)-1 {
			break
		}
		next := sysmbols[rune(s[i+1])]
		cur := sysmbols[rune(s[i])]
		if cur >= next { // 12 写做 XII
			sum += cur
		} else { // 小的数据在大的左边
			sum -= cur
		}
	}
	sum = sum + sysmbols[rune(s[len(s)-1])]
	return sum
}

// longestCommonPrefix 最长公共前缀
func longestCommonPrefix(strs []string) string {
	var result string
	for i := 0; i < len(strs[0]); i++ {
		prev := strs[0][i]
		// 检查当前所有字段的前缀是否一致
		for j := 0; j < len(strs); j++ {
			if i+1 > len(strs[j]) { //检查每个字符串长度
				return result
			}
			cur := strs[j][i]
			if prev != cur {
				return result
			}
		}
		result = result + string(prev)
	}
	return result
}
