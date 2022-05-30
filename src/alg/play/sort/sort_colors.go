package sort

/***

Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
Example 2:

Input: nums = [2,0,1]
Output: [0,1,2]
Example 3:

Input: nums = [0]
Output: [0]

**/

type Record1 struct {
	Key     int
	Counter int
}

func sortColors1(nums []int) {
	// inser sort
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums) && j-1 > 0; j-- {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			} else {
				break
			}
		}
	}
	var counterNums []int
	var recordMap = make(map[int]Record1)
	for _, i := range nums {
		if v, ok := recordMap[i]; ok {
			recordMap[i] = Record1{Key: i, Counter: v.Counter + 1}
		} else {
			recordMap[i] = Record1{Key: i, Counter: 1}
			counterNums = append(counterNums, i)
		}
	}
	var result []int
	for _, i := range counterNums {
		v := recordMap[i]
		var tmp []int
		for j := 0; i < v.Counter; j++ {
			tmp = append(tmp, v.Key)
		}
		result = append(result, tmp...)
	}
	nums = result
}

//
func sortColors2(nums []int) {
	if nums == nil || len(nums) == 1 {
		return
	}
	pivot := 1
	low := 0
	high := len(nums) - 1
	idx := 0
	// move all the elements less than 1 to left, greater than 1 to right
	for idx < high {
		if nums[idx] < pivot {
			swap(nums, idx, low)
			low++
			idx++
		} else if nums[idx] > pivot { // if elemnt is 2, move to right
			swap(nums, idx, high)
			high--
		} else { // if element is 1(pivot), just continue
			idx++
		}
	}
}

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

// insert sort
func sortColors(nums []int) {
	if nums == nil || len(nums) == 1 {
		return
	}
	//sort records accord ascall code value
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums) && j-1 >= 0; j-- {
			if nums[j-1] < nums[j] {
				break // sorted
			} else {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
}
