package dynamic_programming

// 最小下降和，向下、向左下、向右下
// 状态： 当前的和。 选择： 3个移动方向
// dp[i][j]: 从i到J 的最短路径和
func minFallingPathSum(martix [][]int) int {
	var dpTable [][]int
	for i := 0; i < len(martix); i++ {
		for j := 0; j < len(martix[j]); j++ {
			dpTable[i][j] = max4(
				dpTable[i+1][j],   //下
				dpTable[i+1][j+1], // 右下
				dpTable[i-1][j-1], // 作下
			)
		}
	}

	result := 0
	for i := 0; i < len(dpTable); i++ {
		result = max(result, dpTable[len(martix)-1][i])
	}
	return result

}

func max4(a, b, d int) int {
	if a < b {
		a = b
	}
	if a > d {
		return a
	}

	return d
}

// 编辑距离
// 给定2个字符串，计算将S1, 转换成S2的最少操作次数，你可以执行
// 1. 删除一个字符
// 2. 插入一个字符
// 3. 替换一个字符
/**
状态： dp[i][j]: 从S1[i], S2[j]的最小操作次数
*/

func minDistance(s1 []int, s2 []int) int {
	return dpT(s1, s2, len(s1), len(s2))
}

func dpT(s1, s2 []int, i, j int) int {
	if s1[i] == s2[j] {
		return dpT(s1, s2, i-1, j-1)
	}
	a := dpT(s1, s2, i, j-1)   // 删除
	b := dpT(s1, s2, i, j+1)   // 插入一个与j相同的字符
	c := dpT(s1, s2, i-1, j-1) // 替换一个字符
	if a < b {
		a = b
	}
	if a > c {
		return a
	}
	return c
}
