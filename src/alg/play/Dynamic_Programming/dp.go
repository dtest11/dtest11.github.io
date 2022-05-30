package dp

func longestValidParentheses(s string) int {
	panic("longestValidParentheses")
}

func maxSubArray(nums []int) int {
	dpTable := []int{}
	for i := range nums {
		dpTable[i] = nums[0]
	}
	resp := 0
	for n := range nums {
		if dpTable[n-1] > 0 {
			dpTable[n] = dpTable[n-1] + nums[n]
		} else {
			dpTable[n] = nums[n]
		}
		if resp < dpTable[n] {
			resp = dpTable[n]
		}
	}
	return resp
}
