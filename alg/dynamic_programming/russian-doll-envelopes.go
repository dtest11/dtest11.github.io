package dynamic_programming

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	}

	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	// 对于h 求最长递增子序列(这里的序列，可以不连续)
	dp := make([]int, len(envelopes))
	dp[0] = 1
	res := dp[0]
	for i := 1; i < len(envelopes); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] {
				dp[i] = MaxInt(dp[i], dp[j]+1)
			}
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}

func maxEnvelopesV2(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	}

	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	// 对于h 求最长递增子序列(这里的序列，可以不连续)
	dp := make([]int, len(envelopes))
	dp[0] = 1
	var h []int
	for _, v := range envelopes {
		h = append(h, v[1])
	}

	return lengthOfLISV2(h)
}

// 使用二分查找: 最长递增子序列
// 记录每个堆中的数据
func lengthOfLISV2(nums []int) int {
	count := 0                    // 牌堆
	top := make([]int, len(nums)) // 代表每个牌堆上面的数据
	for _, v := range nums {
		// 待处理的牌
		poker := v
		// 找到这张牌 应该放的位置，他要放在某个堆的话，这个堆最上面的数据，必须要比他小，而且这些堆是有序，可以使用二分查找
		// 找到大于poker的第一个
		l := 0
		r := count
		for l < r {
			mid := (l + r) / 2
			if top[mid] < poker {
				l = mid + 1
			} else if top[mid] > poker {
				r = mid
			} else { // top[mid]== poker
				r = mid
			}
		}
		if l == count {
			count = count + 1
		}
		top[l] = poker
	}
	return count
}
