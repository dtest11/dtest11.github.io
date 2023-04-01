package twopointers

import "fmt"

/**
https://leetcode.com/problems/partition-labels/description/
函数结束条件未L<len:
1. 先找到第一个元素的，结束位置[l,r]:对于第一次循环:0..r
2. 检查l+1..r 中的每个元素的最后出现位置，是否结束位置都在r内
2.1 在r 内，找到一个合法的字串，继续，排查第l=r+1,r=dict[l]个字符，用同样的方法
2.2 发现l+1..r中的某个字符的结束位置end,end>r ,那么，移动r的位置到end:这里可以是l+1,r中所有字符的最大end,
r=end ,处理完成之后，l=r+1, r=dict[l]
*
*/

func partitionLabelsV(s string) []int {
	dict := make(map[uint8]int) //  存储某个字母出现的最大索引
	for i, v := range s {
		dict[uint8(v)] = i
	}
	n := len(s)
	// 当前片段开始位置和结束位置
	l := 0
	r := 0
	var result []int
	for l < n { // 这里注意从LK开始作为条件
		r = dict[s[l]]
		j := l + 1
		for j < r { // 检查l+1..r 中的每个元素的最后出现位置，是否在r 内
			temp := dict[s[j]]
			if temp > r { // 当前字符的结束位置超过right
				r = temp // 更新结束的位置
			}
			j = j + 1
		}

		//fmt.Println(s[l:r+1], l, r)
		result = append(result, r+1-l)
		// 重新赋值，找到截至的位置
		// 寻找到合格的一组字符串,移动位置l=r+1,r=dict[l]
		// r 已经计算过，需要r+1
		l = r + 1
		fmt.Println(l, n)
	}
	return result
}

func partitionLabels(s string) []int {
	dict := make(map[uint8]int) //  存储某个字母出现的最大索引
	for i, v := range s {
		dict[uint8(v)] = i
	}
	var result []int
	n := len(s)
	// 当前片段开始位置和结束位置
	l := 0
	r := dict[s[l]]

	for r < n { // 这里注意从r开始作为条件
		j := l + 1
		for j < r && j < n { // 检查l+1..r 中的每个元素的最后出现位置，是否在r 内
			temp := dict[s[j]]
			if temp > r { // 当前字符的结束位置超过right
				r = temp // 更新结束的位置
			}
			j = j + 1
		}

		//fmt.Println(s[l:r+1], l, r)
		result = append(result, r+1-l)
		// 重新赋值，找到截至的位置
		// 寻找到合格的一组字符串,移动位置l=r+1,r=dict[l]
		// r 已经计算过，需要r+1
		l = r + 1
		if l == n {
			break
		}
		r = dict[s[l]]
	}
	return result
}
