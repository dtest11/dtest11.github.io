package removenodesfromlinkedlist

import (
	"testing"
)

func Test_removeZeroSumSublists(t *testing.T) {
	l := &ListNode{Val: 1, Next: &ListNode{Val: -1}}
	removeZeroSumSublists(l)
}
