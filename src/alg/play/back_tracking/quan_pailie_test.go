package back_tracking

import (
	"container/list"
	"fmt"
	"testing"
)

func Test_permute(t *testing.T) {
	result := permute([]int{1, 2, 3, 4}) // n阶乘 1*2 *3 *4
	t.Logf("%+v\n", result)
	t.Log(len(result))
}

func Test_permute1(t *testing.T) {
	// create a new link list
	alist := list.New()

	fmt.Println("Size before : ", alist.Len()) // list size before

	// push element into list
	alist.PushBack("a")
	alist.PushBack("b")

	
	alist.PushBack("c")

	fmt.Println("Size after insert(push): ", alist.Len()) // list size after

	// iterate over list elements
	for e := alist.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(string))
	}
}
