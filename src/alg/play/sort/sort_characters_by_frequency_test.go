package sort

import "testing"

func Test_frequencySort(t *testing.T) {
	result := frequencySort("dacccccb")
	t.Log(result) // cccccabd

}
