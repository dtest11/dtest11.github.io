package array

import "testing"

func Test_twoSum(t *testing.T) {
	data := []int{2, 7, 11, 15}
	target := 9
	t.Log(twoSum(data, target))
	data = []int{3, 2, 4}
	target = 6
	t.Log(twoSum(data, target))
}
