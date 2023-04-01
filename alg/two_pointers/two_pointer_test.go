package twopointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	var str = "ababcbacadefegdehijhklij"
	res := partitionLabels(str)
	assert.Equal(t, []int{9, 7, 8}, res)
	str = "qiejxqfnqceocmy"
	res = partitionLabels(str)
	assert.Equal(t, []int{13, 1, 1}, res)
}

func Test_latestTimeCatchTheBus(t *testing.T) {
	type args struct {
		buses      []int
		passengers []int
		capacity   int
		expect     int
	}

	tests := []args{
		{
			buses:      []int{10, 20},
			passengers: []int{2, 17, 18, 19},
			capacity:   2,
			expect:     16,
		},

		{
			buses:      []int{20, 30, 10},
			passengers: []int{19, 13, 26, 4, 25, 11, 21},
			capacity:   2,
			expect:     20,
		},

		{
			buses:      []int{3},
			passengers: []int{2, 3},
			capacity:   2,
			expect:     1,
		},
		{
			buses:      []int{2},
			passengers: []int{2},
			capacity:   1,
			expect:     1,
		},
		{
			buses:      []int{2},
			passengers: []int{2},
			capacity:   2,
			expect:     1,
		},
	}

	for _, tt := range tests {
		res := latestTimeCatchTheBus(tt.buses, tt.passengers, tt.capacity)
		assert.Equal(t, tt.expect, res)
	}
}
