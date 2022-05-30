package string

import "testing"

/**

Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.



Example 1:

Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
Example 2:

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9


Constraints:

0 <= nums.length <= 105
-109 <= nums[i] <= 109
**/

// 使用hash 字典
func longestConsecutive(nums []int) int {
	var record = make(map[int]bool)

	for _, v := range nums {
		record[v] = true
	}
	maxN := 0
	for k, v := range record {
		if !v {
			continue
		}
		currentL := 1
		left, right := k-1, k+1
		for {
			v1, ok := record[left]
			if ok && v1 {
				record[left] = false
				left--
				currentL++
			} else {
				break
			}
		}

		for {
			v1, ok := record[right]
			if ok && v1 {
				record[right] = false
				currentL++
				right++
			} else {
				break
			}
		}
		if maxN < currentL {
			maxN = currentL
		}

	}
	return maxN
}

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "search", args: args{nums: []int{100, 4, 200, 1, 3, 2}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("%s  got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
