package dynamic_programming

import "github.com/dtest11/alg/tree"

// 相邻的屋子不能抢
func rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}

	var dp = make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < l; i++ {
		if i == 0 {
			dp[0][0] = 0
			dp[0][1] = nums[0]
		} else {
			dp[i][0] = max(dp[i-1][1], dp[i-1][0]) // 当前不抢
			dp[i][1] = nums[i] + dp[i-1][0]
		}
	}
	return max(dp[l-1][0], dp[l-1][1])
}

// / 头部和尾部不能同时被抢
func robCircle(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	i := rob(nums[:l-1])
	j := rob(nums[1:])
	return max(i, j)
}

type TreeNode = tree.TreeNode

type cache struct {
	rob  map[*TreeNode]int
	skip map[*TreeNode]int
}

func newCache() *cache {
	return &cache{rob: map[*TreeNode]int{}, skip: map[*TreeNode]int{}}
}

func robTree(root *TreeNode) int {
	c := newCache()
	return robTreeHelp(root, c)
}

func robTreeHelp(root *TreeNode, c *cache) int {
	if root == nil {
		return 0
	}
	// check cache
	a, ok1 := c.rob[root]
	b, ok2 := c.skip[root]
	if ok1 && ok2 {
		if a > b {
			return a
		} else {
			return b
		}
	}

	//抢,然后去下下家
	temp := root.Val
	if root.Left != nil {
		temp += robTreeHelp(root.Left.Left, c)
		temp += robTreeHelp(root.Left.Right, c)
	}
	if root.Right != nil {
		temp += robTreeHelp(root.Right.Right, c)
		temp += robTreeHelp(root.Right.Left, c)
	}
	c.rob[root] = temp

	// 不抢,直接去下家
	c.skip[root] = robTreeHelp(root.Left, c) + robTreeHelp(root.Right, c)
	// 2个情况取最大值
	a = c.rob[root]
	b = c.skip[root]
	if a > b {
		return a
	} else {
		return b
	}
}
