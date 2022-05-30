package leetcode

import "golang.org/x/exp/constraints"

type PriorityQueue[T constraints.Ordered] struct {
	Val T
}
