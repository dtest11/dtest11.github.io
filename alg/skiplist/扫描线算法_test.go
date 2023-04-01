package skiplist

import (
	"testing"
)

func TestSortNode(t *testing.T) {
	_ = []Node{
		{0, 30},
		{50, 10},
		{15, 20},
	}
	//sort.Slice(nums, func(i, j int) bool {
	//	if nums[i].start == nums[j].start {
	//		return nums[i].end < nums[j].end
	//	}
	//	return nums[i].start < nums[j].start
	//})
	//t.Log(nums)
	var payload = [][2]int{{0, 30}, {50, 10}, {15, 20}}

	t.Log(minMeetingRooms(payload))
}
