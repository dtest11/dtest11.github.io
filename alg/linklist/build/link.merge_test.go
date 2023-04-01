package build

import (
	"fmt"
	"testing"
)

func Test_mergeTwoLink(t *testing.T) {
	var tmp []*link[int32]
	fmt.Println(tmp)
	first := NewNode(0)
	second := NewNode(100)
	first = mergeTwoLink(first, second, -1)
	for first != nil {
		fmt.Println(first.data)
		first = first.Next
	}
}

func generateLink() *link[int] {
	root := NewNode(0)
	head := root
	for i := 1; i < 10; i++ {
		root.Next = NewNode(i)
		root = root.Next
	}
	return head
}

func TestLastKNode(t *testing.T) {
	node := generateLink()
	result := lastKNode(node, 5)
	t.Log(result.data)
}
