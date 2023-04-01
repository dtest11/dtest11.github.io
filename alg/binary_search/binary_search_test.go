package binarysearch

import "testing"

func Test_shipWithinDays(t *testing.T) {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	days := 5
	t.Log(shipWithinDays(weights, days))
}
