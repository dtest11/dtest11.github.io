package reverseBetween

import (
	"fmt"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	left := 2
	right := 4
	fmt.Println(a[:left-1])
	fmt.Println(a[left-1 : right])
	fmt.Println(a[right:])

	a = []int{3, 5}
	b := &ListNode{Val: 3, Next: &ListNode{Val: 5}}
	reverseBetween(b, 1, 2)
}
