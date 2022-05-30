package bubble

import "testing"

func TestBubble(t *testing.T) {
	data := []int{10, 4, 2, 3, 7, 9, 1}
	Bubble(data)
	t.Log("hello")
	t.Log(data)
}
