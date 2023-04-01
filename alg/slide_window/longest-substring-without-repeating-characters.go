package slidewindow

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

//输入一个字符串S[cabd],另一个T[ab]， 找到T的全排列的
/**
Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.

A substring is a contiguous sequence of characters within the string.
Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
Example 2:

Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
Example 3:

Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.

**/

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)
	validSize := 0
	for i := range t {
		v := t[i]
		val, ok := need[v]
		if ok {
			need[v] = val + 1
		} else {
			need[v] = 1
		}
	}

	left := 0
	right := 0
	start := 0
	end := 0
	for right < len(s) {
		cur := s[right]
		if _, ok := need[cur]; ok {
			window[cur] += 1
			if window[cur] == need[cur] {
				validSize++
			}
		}
		right++
		// 检查是否要shrink
		for validSize == len(need) {
			if right-left < (end - start) {
				start = left
				end = right
			}
			cur := s[left]
			if _, ok := need[cur]; ok {
				window[cur]--
				validSize--
			}
			left++
		}
	}
	return s[start:end]
}

/**

LeetCode 567 题，Permutation in String，难度 Medium：
给定2个字符串S1,S2，写一个函数判断S2是否是包含S1的排列


Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.


Example 1:

Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:

Input: s1 = "ab", s2 = "eidboaoo"
Output: false

**/

func checkInclusion(s1 string, s2 string) bool {
	needCount := make(map[byte]int)
	window := make(map[byte]int)
	for i := range s1 {
		needCount[s1[i]]++
	}
	left, right, valid := 0, 0, 0
	for right < len(s2) {
		cur := s2[right]
		right++

		if needCount[cur] != 0 {
			window[cur]++
			if window[cur] == needCount[cur] {
				valid++
			}
		}
		// 检查窗口是否收缩
		for right-left >= len(s1) {
			if valid == len(needCount) {
				return true
			}
			cur := s2[left]
			left++
			if needCount[cur] != 0 {
				if needCount[cur] == window[cur] {
					valid--
				}
				window[cur]--
			}
		}
	}
	return false
}
