package main

import (
	"testing"
)

func Test_permute(t *testing.T) {
	var input = []int{1, 2, 3}
	result := permute[int](input)
	for _, v := range result {
		t.Log(v)
	}
}
