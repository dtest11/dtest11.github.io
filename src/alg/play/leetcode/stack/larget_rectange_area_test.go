package stack

import (
	"fmt"
	"testing"
)

func largestRectangleArea(heights []int) int {
	var newHeights = make([]int, len(heights)+2)
	newHeights[0] = 0
	newHeights[len(newHeights)-1] = 0
	for i := 1; i < len(heights)+1; i++ {
		newHeights[i] = heights[i-1]
	}
	var st = NewStack()
	var result int
	for i := 0; i < len(newHeights); i++ {
		// 如果栈不为空，并且当前的元素 小于 栈顶元素，那么以栈顶元素为高的矩形的面积可以确定
		for st.Len() > 0 && newHeights[i] < newHeights[st.Peek().(int)] {
			cur := st.Pop().(int) // 弹出栈顶元素
			curHeight := newHeights[cur]
			// 栈顶元素弹出之后，新的栈顶元素就是它的左边界
			leftIndex := st.Peek().(int)
			// 右边界就是当前考察的i
			rightIndex := i
			// 计算矩形的面积
			curWidth := rightIndex - leftIndex - 1
			tmp := curWidth * curHeight
			if tmp > result {
				result = tmp
			}
		}
		// 当前考察的元素入栈
		st.Push(i)
	}
	return result
}

func largestRectangleAreaB(heights []int) int {
	lg := len(heights)
	var result int
	for i := 0; i < lg; i++ {
		cur := heights[i] // 固定高度
		w := 0
		for j := 0; j < lg; j++ {
			next := heights[j]
			if next >= cur {
				w++
				continue
			} else {
				t := w * cur
				if t > result {
					result = t
				}
				w = 0
			}
		}
		t := w * cur
		if t > result {
			result = t
		}
	}
	return result
}

func Test_lar(t *testing.T) {
	var data = []int{2, 4}
	fmt.Println(largestRectangleArea(data))
	data = []int{2, 0, 2}
	fmt.Println(largestRectangleArea(data))
	data = []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(data))

}
