package quick

import (
	"fmt"
	"testing"
)

func Test_quick(t *testing.T) {
	var nums = []int{3, 2, 4, 1}
	quick(nums, 0, len(nums)-1)
	fmt.Println(nums)
	quick(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
