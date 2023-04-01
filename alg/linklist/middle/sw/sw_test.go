package sw

import "testing"

func Test_xxx(t *testing.T) {
	l := &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4}}}}
	reorderList(l)
}
