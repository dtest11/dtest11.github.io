package palindrome_linklist

import (
	"fmt"
	"testing"
)

func Test_getMiddle(t *testing.T) {
	var list = &LinkList{}
	for i := 0; i < 11; i++ {
		list.InsertFront(i)
	}
	list.Print()
	head := middleNode(list.head)
	t.Log(head.Val)

	fmt.Println(2 << 2)
}
