package back_tracking

/**
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回 s 所有可能的分割方案。

示例:

输入:"aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]

https://leetcode.com/problems/palindrome-partitioning/

131. 分割回文串

回溯算法  back_tracking
**/
// TODO
func partition(s string) [][]string {
	var result [][]string
	tmp := make([]string, len(s))
	partitionBackTracking(s, 0, tmp, &result)
	return result
}

func partitionBackTracking(s string, startIndex int, tmp []string, result *[][]string) {
	if startIndex >= len(s) { // 如果起始位置已经大于s的大小，说明已经找到了一组分割方案了
		*result = append(*result, tmp)
		return
	}
	for i, _ := range s {
		if !isPalindrome(s, startIndex, i) {
			continue
		}
		str := s[startIndex : i-startIndex+1] // / 获取[startIndex,i]在s中的子串
		tmp = append(tmp, str)
		partitionBackTracking(s, startIndex+1, tmp, result)
		tmp = tmp[:len(tmp)-1]
	}
}

func isPalindrome(s string, start, end int) bool {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

/**
void backtracking(参数) {
    if (终止条件) {
        存放结果;
        return;
    }

    for (选择：本层集合中元素（树中节点孩子的数量就是集合的大小）) {
        处理节点;
        backtracking(路径，选择列表); // 递归
        回溯，撤销处理结果
    }
}


*/
