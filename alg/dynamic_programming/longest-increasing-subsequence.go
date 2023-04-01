package dynamic_programming

// lengthOfLIS 动态规划解法
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = MaxInt(dp[i], dp[j]+1)
			}
		}
	}
	var res = 0
	for _, v := range dp {
		if res < v {
			res = v
		}
	}
	return res
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS_binary_search(nums []int) int {
	var piles [][]int // 堆
	for _, v := range nums {
		// v 为当前需要处理的数据
		// 由于堆从左边到最后一个堆，数据是单调递增的，可以使用二分查找法
		l := 0
		r := len(piles)
		for l < r {
			mid := (l + r) / 2
			pile := piles[mid]
			topVal := pile[len(pile)-1]
			if topVal > v {
				r = mid
			} else if topVal < v {
				l = mid + 1
			} else { // topVal==v
				r = mid
			}
		}
		if l <= len(piles)-1 {
			piles[l] = append(piles[l], v)
		} else {
			// 查找结束，检查是否寻找到合适的堆
			piles = append(piles, []int{v})
		}
	}
	return len(piles)
}

// 返回字典序最小的递增序列
func lengthOfLIS_min(nums []int) []int {
	maxLen := 1
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		if maxLen < dp[i] {
			maxLen = dp[i]
		}
	}
	var ans = make([]int, maxLen)
	for i := len(nums) - 1; i >= 0; i-- {
		if dp[i] == maxLen {
			ans[i] = nums[i]
			maxLen = maxLen - 1
		}
	}
	return ans
}
