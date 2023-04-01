package skiplist

import (
	"fmt"
	"sort"
)

/**
https://www.jiuzhang.com/solution/meeting-rooms-ii/
**/

type NodeVal struct {
	time int
	cost int
}

func minMeetingRooms(intervals [][2]int) int {
	var roots []NodeVal
	for _, v := range intervals {
		roots = append(roots, NodeVal{time: v[0], cost: 1})
		roots = append(roots, NodeVal{time: v[1], cost: -1})
	}
	sort.Slice(roots, func(i, j int) bool {
		if roots[i].time == roots[j].time {
			return roots[i].cost < roots[j].cost
		}
		return roots[i].time < roots[j].time
	})
	result := 0
	tmp := 0
	fmt.Println(roots)
	for i := 0; i < len(roots); i++ {
		tmp += roots[i].cost

		if tmp > result {
			result = tmp
		}
	}
	return result
}
