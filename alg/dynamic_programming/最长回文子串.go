package dynamic_programming

// longestPalindrome 返回给定字符串 s 的最长回文子串的长度
func longestPalindrome(s string) string {
	n := len(s)

	if n <= 1 {
		return s
	}
	// fast path
	if n == 2 {
		if s[0] == s[1] {
			return s
		}
		return ""
	}

	// dp[i][j] 表示 s[i..j] 是否回文，j >= i
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true // 易知，单个字符 s[i..i] 构成回文
	}

	// 记录最大回文子串的长度，至少为 1
	maxLength := 1
	var left int
	var right int

	// 考虑递推
	// 主要的递推关系是 dp[i][j] = dp[i+1][j-1]
	// 所以倒序遍历 i ，才可以形成递推
	// 从下到上，从左到右遍历，这样保证dp[i + 1][j - 1]都是经过计算的
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-1 >= i+1 { // 子串 s[i+1..j-1] 有效性
					if dp[i+1][j-1] {
						dp[i][j] = true
					}
				} else {
					// 此时 j < i + 2 即 j <= i+1
					// 再之 s[i] == s[j]，必回文
					dp[i][j] = true
				}
			}

			if dp[i][j] {
				// 更新最大长度
				length := j - i + 1
				if length > maxLength {
					maxLength = length
					left = i
					right = j
				}
			}
		}
	}

	return s[left : right+1]
}
