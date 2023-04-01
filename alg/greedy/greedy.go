package greedy

import "sort"

// intervalSchedule 不相交的区间
func intervalSchedule(data [][2]int) int {
	sort.Slice(data, func(i, j int) bool {
		return data[i][1] < data[j][1]
	})
	lastEnd := data[0][1]
	res := 0
	for i := 1; i < len(data)-1; i++ {
		cur := data[i][0]
		if lastEnd < cur {
			res++
			lastEnd = data[i][1]
		}
	}
	return res
}

func minMeetingRooms(meetings [][2]int) int {
	var start []int
	var end []int
	for _, v := range meetings {
		start = append(start, v[0])
		end = append(start, v[1])
	}
	sort.Slice(start, func(i, j int) bool {
		return start[i] < start[j]
	})
	sort.Slice(end, func(i, j int) bool {
		return end[i] < end[j]
	})
	startIndex := 0
	endIndex := 0
	count := 0
	for startIndex < len(meetings) && endIndex < len(meetings) {
		if start[startIndex] < end[endIndex] {
			startIndex++
			count++
		} else {
			endIndex++
			count--
		}
	}
	return count
}

// https://leetcode.jp/leetcode-45-jump-game-ii-%E8%A7%A3%E9%A2%98%E6%80%9D%E8%B7%AF%E5%88%86%E6%9E%90/
// data[i]表示最多可以调n步，问跳到末尾需要的最少步数是多少
func jumpII(data []int) int {
	n := len(data)
	index := 0
	res := 0
	for index <= n-1 {
		// 当前节点能跳到的left,right
		left := index + 1
		right := index + data[index]
		max := 0 // 所有跳跃距离中，最远的
		// 查看范围中的每一个点
		if right >= n-1 {
			return res + 1
		}
		for i := left; i <= right; i++ {
			cur := i + data[i] // 当前距离加上能跳远的距离
			if cur >= max {
				max = cur
				index = i // 下一跳的距离
			}
		}
		res++
	}
	return res
}

// 是否能跳到末尾
func jump(nums []int) bool {
	max := 0
	for i := 0; i < len(nums); i++ {
		cur := nums[i] + i
		if max > cur {
			max = cur
		}
	}
	return max > n-1
}
