package removeDuplicateNodes

import "testing"

// [1, 2, 3, 3, 2, 1
func Test_remove(t *testing.T) {
	var l = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
		},
	}
	res := removeDuplicateNodes(l)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}
