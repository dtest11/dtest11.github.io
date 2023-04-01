package convert_binary_number

import "testing"

func Test_getDecimalValue(t *testing.T) {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	val := getDecimalValue(l)
	t.Log(val)
}
