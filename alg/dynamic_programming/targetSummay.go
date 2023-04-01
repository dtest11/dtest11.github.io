package dynamic_programming

/*
*
给你一个整数数组 nums 和一个整数 target 。

向数组中的每个整数前添'+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：

例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/target-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
示例 2：

输入：nums = [1], target = 1
输出：1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/target-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

	背包问题
*/
func findTargetSummaryWays(vec []int, target int) int {
	var dpTable [][]int
	for i := 1; i < len(vec); i++ {
		for j := 0; j < target; j++ {
			if j >= vec[i-1] {
				// 两种选择的结果之和
				dpTable[i][j] = dpTable[i-1][j] + dpTable[i-1][j-vec[i-1]]
			} else {
				// 背包的空间不足，只能选择不装物品 i
				dpTable[i][j] = dpTable[i-1][j]
			}
		}
	}
	return dpTable[len(vec)][target]
}

/**
状态：当前的目标和
选择： 使用+1, 或者使用-1
dp[i]: 从0-到i的最大目标和
*/

/**
回溯算法：
for i in 选择列表
   back.push(i)
   backtrack(back, 选择列表[i..end])
   back.pop()
*/

func backtrace(vec []int, target int) int {
	var result int
	backTrackImpl(vec, 0, 0, target, &result)
	return result
}

func backTrackImpl(vec []int, startIndex int, currentVal int, target int, result *int) {
	if currentVal == target {
		*result++
	}
	if startIndex == len(vec) {
		return
	}
	for i := startIndex; i < len(vec); i++ {
		currentVal += vec[i]
		backTrackImpl(vec, startIndex+1, currentVal, target, result)
		currentVal -= vec[i]
	}
}
