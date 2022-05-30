package back_tracking

import (
	"container/list"
	"fmt"
)

func permute(nums []int) [][]int {
	var track = list.New()
	var result1 [][]int
	var permuteBackTrack = func(nums []int, track *list.List) {}
	permuteBackTrack = func(nums []int, track *list.List) {
		////路径：记录track
		////选择列表： nums中不在于tack 中的元素
		////结束条件：nums中的元素全部出现在track中
		if track.Len() == len(nums) {
			ele := LoopElement(track)
			result1 = append(result1, ele)
			return
		}

		for _, v := range nums {
			if isContain(track, v) {
				continue
			}
			tmp := track.PushBack(v)      // 做选择
			permuteBackTrack(nums, track) // 进入下一层决策树
			track.Remove(tmp)             // undo
		}
	}
	permuteBackTrack(nums, track)
	return result1
}

func isContain(l *list.List, element int) bool {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == element {
			return true
		}
	}
	return false
}

func LoopElement(l *list.List) (result []int) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
		result = append(result, e.Value.(int))
	}
	return
}

// 第二个解法不用 linkList

func permute2(nums []int) [][]int {
	var ret [][]int
	permuteBackTrack2(nums, make([]bool, len(nums)), []int{}, &ret)
	return ret
}

func permuteBackTrack2(nums []int, visited []bool, temp []int, ret *[][]int) {
	if len(temp) == len(nums) {
		*ret = append(*ret, temp)
	}
	for i, n := range nums {
		if visited[i] {
			continue
		}
		visited[i] = true
		permuteBackTrack2(nums, visited, append(temp, n), ret)
		visited[i] = false
	}

}
