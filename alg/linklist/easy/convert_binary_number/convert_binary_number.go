package convert_binary_number

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 101 => 5
// 2进制转换成10 进制
// 注意 2的次方 => 1<< n
func getDecimalValueV2(head *ListNode) int {
	// reverse
	var cur = head
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	var result int
	var index int
	for prev != nil {
		if prev.Val != 0 {
			result += (1 << index)
		}
		prev = prev.Next
		index++
	}

	return result
}

func getDecimalValue(head *ListNode) int {
	return getDecimalValueV2(head)
	var tmp []int
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	// reverse
	var rTmp []int
	for i := len(tmp) - 1; i >= 0; i-- {
		rTmp = append(rTmp, tmp[i])
	}
	var result []int
	for i := 0; i < len(rTmp); i++ {
		val := rTmp[i]
		if val != 0 {
			fmt.Println(1 << i)
			result = append(result, 1<<i)
		}
	}
	// sum res
	var res int
	for i := 0; i < len(result); i++ {
		res += result[i]
	}
	return res
}
