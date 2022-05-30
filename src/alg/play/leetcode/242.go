package leetcode

import (
	"fmt"
	"strings"
)

/*****

该题意是比较两个字符串s和t，其中t是次序打乱的字符串，如果两个字符串相同则返回true，否则false
****/

func isAnagram(s string, t string) bool {
	ss := strings.Split(s, "")
	ts := strings.Split(t, "")

	if len(ss) != len(ts) {
		return false
	}

	record := make(map[string]int)

	for _, val := range ss {
		if old, ok := record[val]; ok {
			record[val] = old + 1
		} else {
			record[val] = 1

		}
	}
	for _, val := range ts {
		old, ok := record[val]
		if !ok {
			return false
		}

		if old < 0 {
			return false
		}
		if old-1 == 0 {
			delete(record, val)
		} else {
			record[val] = old - 1
		}
	}
	counter := 0
	for _, val := range record {
		counter += val
	}
	if counter > 0 {
		return false
	}
	return true
}

func main() {
	a := "anagram"
	b := "nagaram"

	//a="aacc"
	//b="ccac"

	result := isAnagram(a, b)
	fmt.Println(result)
}
