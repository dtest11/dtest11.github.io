package sort

import "testing"

func Test_sortColors(t *testing.T) {
	input := []int{2, 0, 2, 1, 1, 0}
	sortColors(input)
	t.Log(input)
}
