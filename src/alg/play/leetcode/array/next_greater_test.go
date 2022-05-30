package array

import (
	"fmt"
	"testing"
)

// 单调栈的使用场景
/**
根据题意，数组 nums1 视为询问。我们可以：

1.先对 nums2 中的每一个元素，求出它的右边第一个更大的元素；
2.将上一步的对应关系放入哈希表（HashMap）中；
3.再遍历数组 nums1，根据哈希表找出答案。

作者：LeetCode
链接：https://leetcode-cn.com/problems/next-greater-element-i/solution/xia-yi-ge-geng-da-yuan-su-i-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
**/

// nums1 是 nums2的子集，并且没有重复元素，那么就从nums2 先解决问题，解决之后，再把答案应用到 nums1
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	record := map[int]int{}
	var stack gStack
	for i := 0; i < len(nums2); i++ {
		cur := nums2[i]
		// 首先需要判断栈顶元素，跟当前元素的大小关系
		// 1. stack.Top() > cur  符合单调栈的
		// 2. stack.Top() > Cur 不符合单调栈的递增关系，需要将栈顶元素，pop移除，直到找到
		for stack.NotEmpty() && stack.Top() < cur {
			// 移除栈顶的元素，一直到栈顶元素布 大于cur 或者 stack.empty
			e, _ := stack.Pop()
			record[e] = cur
		}
		stack.Push(cur)
	}

	var result []int
	for _, e := range nums1 {
		v, ok := record[e]
		if ok {
			result = append(result, v)
		} else {
			result = append(result, -1)
		}
	}
	return result
}

type gStack struct {
	length int
	data   []int
}

func (s *gStack) Push(ele int) {
	s.length++
	s.data = append(s.data, ele)
}

func (s *gStack) Top() int {
	return s.data[s.length-1]
}

func (s *gStack) NotEmpty() bool {
	return s.length > 0
}

// Pop 移除栈顶元素
func (s *gStack) Pop() (int, bool) {
	if s.length == 0 {
		return 0, false
	}
	index := len(s.data) - 1 // Get the index of the top most element.
	element := s.data[index] // Index into the slice and obtain the element.
	s.data = s.data[:index]  // Remove it from the stack by slicing it off.
	s.length--
	return element, true
}

func Test_nextGreaterElement(t *testing.T) {
	a := []int{4, 1, 2}
	b := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(a, b))
}
