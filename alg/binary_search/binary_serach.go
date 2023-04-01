package binarysearch

import (
	"math"
)

// O(log n)。
func binarySeach(nums []int, target int) {
	l := 0
	r := len(nums)
	for l < r {
		mid := (r-l)/2 + l
		cur := nums[mid]
		if cur == target {
			return
		} else if cur < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
}

// lower_bound 搜索左侧边界
func lower_bound(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		mid := (r-l)/2 + l
		cur := nums[mid]
		if cur < target {
			l = mid + 1
		} else if cur > target {
			mid = r
		} else { // cur==target.收缩右边界
			r = mid
		}
	}
	return l
}

// right_bound 查找右边界
func right_bound(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		mid := (r-l)/2 + l
		cur := nums[mid]
		if cur < target {
			l = mid + 1
		} else if cur > target {
			r = mid
		} else { // cur ==target 收缩左边界
			l = mid + 1
		}
	}
	return l - 1
}

func minEatingSpeed(piles []int, h int) int {
	l := 1                  //代表吃香蕉的速度
	r := int(math.Pow10(9)) // 一堆里面最多的数量
	for l < r {
		mid := (r-l)/2 + l
		hour := getHour(piles, mid)
		if hour < h { // 时间花的太少，要吃得更慢点
			r = mid
		} else if hour > h { // 时间花的太多，要吃得更快点
			l = mid + 1
		} else { // 收缩右边界
			r = mid
		}
	}
	return l
}

func getHour(pile []int, x int) int {
	h := 0
	for _, v := range pile {
		h += v / x
		if v%x != 0 { // 整数时间内不能吃完，还需要一个小时
			h += 1
		}
	}
	return h
}

func shipWithinDays(weights []int, days int) int {
	//确定 最大的运载量：应该是所有货物的总和：一天运完
	// 最小的运载量 ：是货物的最大值 、、货物不可能拆封
	l := 0
	r := 1 //  // 注意，right 是开区间，所以额外加一
	for _, v := range weights {
		if v > l {
			l = v
		}
		r += v
	}

	for l < r {
		mid := (r-l)/2 + l
		cur := getDays(weights, mid)
		if cur < days {
			r = mid
		} else if cur > days {
			l = mid + 1 // 需要让运载能力更大一些
		} else {
			r = mid // 收缩右边界
		}
	}
	return l
}

// X 代表运载能力
/**
由于我们必须按照数组weights中包裹的顺序进行运送，因此我们从数组
weights 的首元素开始遍历，将连续的包裹都安排在同一天进行运送。当这批包裹的重量大于运载能力
x 时，我们就需要将最后一个包裹拿出来，安排在新的一天，并继续往下遍历。当我们遍历完整个数组后，
就得到了最少需要运送的天数。

**/

func getDays(weights []int, maxCapacity int) int {
	day := 1 // 注意初始化为1
	capacity := 0
	for _, v := range weights {
		capacity = capacity + v
		if capacity > maxCapacity {
			day += 1
			capacity = v
		}

	}
	return day
}

func maximizeSweetness(sweetness []int, k int) int {
	panic("not implement")

	// sum := 0
	// for _, v := range sweetness {
	// 	sum += v
	// }
	// l := 0
	// r := sum / k
	// r = r + 1
	// for l < r {
	// 	mid := (r-l)/2 + l
	// 	cur := fmaximizeSweetness(sweetness, k)
	// }
}

func fmaximizeSweetness(sweetness []int, k int, target int) int {
	panic("not implement")
}
