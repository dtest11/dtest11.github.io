package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var data = map[int][]int{
	9: []int{2, 7, 11, 15},
	6: []int{3, 3},
}

type d struct {
	Target int
	Ele    []int
	Want   interface{}
}

func getEle() []d {
	return []d{
		{
			Target: 9,
			Ele:    []int{-1, 0, 1, 2, -1, -4},
			Want:   []int{0, 1},
		},
		{
			Target: 6,
			Ele:    []int{3, 2, 4},
			Want:   []int{1, 2},
		},
		{
			Target: 6,
			Ele:    []int{3, 3},
			Want:   []int{0, 1},
		},
		{
			Target: -8,
			Ele:    []int{-1, -2, -3, -4, -5},
			Want:   []int{2, 4},
		},
	}
}

func get3Sum() []d {
	return []d{
		{
			Target: 0,
			Ele:    []int{-1, 0, 1, 2, -1, -4},
			Want:   [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		//{
		//	Target: 0,
		//	Ele:    []int{},
		//	Want:   [][]int{},
		//},
		//{
		//	Target: 0,
		//	Ele:    []int{0},
		//	Want:   [][]int{},
		//},
		//{
		//	Target: 0,
		//	Ele:    []int{-2, 0, 0, 2, 2},
		//	Want:   [][]int{{-2, 0, 0, 2, 2}},
		//},
		{
			Target: 0,
			Ele: []int{
				-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
			Want: [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}},
		},
	}
}

func get4Sum() []d {
	return []d{
		//{
		//	Target: 0,
		//	Ele:    []int{1, 0, -1, 0, -2, 2},
		//	Want:   [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		//},
		{
			Target: 1,
			Ele:    []int{0, 0, 0, 0},
			Want:   nil,
		},
		{
			Target: 3,
			Ele:    []int{-1, 2, 2, -5, 0, -1, 4},
			Want:   [][]int{{-5, 2, 2, 4}, {-1, 0, 2, 2}},
		},
		//{
		//	Target: 0,
		//	Ele:    []int{-2, 0, 0, 2, 2},
		//	Want:   [][]int{{-2, 0, 0, 2, 2}},
		//},
		//{
		//	Target: 0,
		//	Ele: []int{
		//		-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
		//	Want: [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}},
		//},
	}
}

func TestSort(t *testing.T) {
	for _, v := range data {
		sort(v)
		t.Log(v)
	}
}

func TestTWoSum(t *testing.T) {
	for _, v := range getEle() {
		res := twoSum(v.Ele, v.Target)
		assert.Equal(t, v.Want, res, fmt.Sprintf("%+v,Error", v))
	}
}

func Test3Sum(t *testing.T) {
	for _, v := range get3Sum() {
		res := threeSumTarget(v.Ele, v.Target)
		assert.Equal(t, v.Want, res, fmt.Sprintf("%+v,Error", v))
	}
}

func Test4Sum(t *testing.T) {
	for _, v := range get4Sum() {

		res := fourSumTarget(v.Ele, v.Target)
		assert.Equal(t, v.Want, res, fmt.Sprintf("%+v,Error", v))
	}
}
