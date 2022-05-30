package string

import (
	"fmt"
	"testing"
)

func permute(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{nums}
	}
	var result [][]int
	backtrack(&result,nums,0,len(nums))
	return  result
}

func backtrack(result *[][]int,nums []int, left ,right int)  {
	if left == right{
		t := make([]int, len(nums))
		copy(t,nums)
		*result =append(*result,t)
	}
	for i :=left; i <right;i++{
		// swap  left, i
		if i> left && nums[i]== nums[i-1]{
			continue // duplicate
		}
		swapV2(nums,i,left)
		backtrack(result,nums,left+1,right)
		swapV2(nums,left,i)
	}
}

func swapV2(nums []int, left ,right int)  {
	nums[left], nums[right] = nums[right], nums[left]
}

func Test_permute(t *testing.T)  {
	var source=[]int{1,2,3,4}
	result :=permute(source)
	fmt.Println(len(result))
}

