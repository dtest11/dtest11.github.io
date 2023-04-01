package getKthFromEnd

import (
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	l1 := &ListNode{
		Val: -9,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val:  7,
			Next: nil,
		},
	}
	l := mergeTwoLists(l1, l2)
	for l != nil {
		t.Log(l.Val)
		l = l.Next
	}
}
