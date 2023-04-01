package revese_print

import (
	"testing"
)

func Test_reversePrint(t *testing.T) {
	var l = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	f := reversePrint(l)
	t.Log(f)
}
