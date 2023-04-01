package dynamic_programming

/**
dp[i,j] :s1[0..i],s2[0..j]的最长公共子序列,dp[s1,s2]
**/

func longestCommonSubsequence(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}
	m := len(text1)
	n := len(text2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		s1 := text1[i-1] // i 从1开始，要计算到0位置的字符
		for j := 1; j <= n; j++ {
			s2 := text2[j-1]
			if s1 == s2 {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = MaxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

// 输出具体的字符串:最长字串是什么
func longestCommonSubsequenceV2(text1 string, text2 string) string {
	if text1 == "" || text2 == "" {
		return ""
	}
	m := len(text1)
	n := len(text2)

	dp := make([][]string, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]string, n+1)
	}

	for i := 1; i <= m; i++ {
		s1 := text1[i-1] // i 从1开始，要计算到0位置的字符
		for j := 1; j <= n; j++ {
			s2 := text2[j-1]
			if s1 == s2 {
				dp[i][j] = dp[i-1][j-1] + string(s1)
			} else {

				//dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				dp[i][j] = dp[i-1][j]
				if len(dp[i][j]) < len(dp[i][j-1]) {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[m][n]
}
