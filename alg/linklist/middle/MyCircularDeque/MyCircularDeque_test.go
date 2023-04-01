package mycirculardeque

import (
	"fmt"
	"testing"
)

func Test_linklist_DeleteLast(t *testing.T) {
	list := newLinkList()
	fmt.Println(list.IsEmpty())
	var data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range data {
		//list.InsertFront(v)
		list.InsertLast(v)
	}
	list.Print()

	list.DeleteLast()
	list.Print()

	list.DeleteFront()
	list.Print()
	//
	//list.ReversePrint()

	//for _, v := range data {
	//	list.InsertFront(v)
	//}
	//list.Print()
	//list.ReversePrint()
	m := Constructor(4)
	m.InsertFront(9)
	m.DeleteLast()
	m.GetRear()
	m.GetFront()
	m.DeleteFront()
	m.InsertFront(6)
	m.InsertLast(5)
	m.InsertFront(9)
	m.GetFront()
	m.InsertFront(6)
}

/**
["MyCircularDeque",
"insertFront",
"deleteLast",
"getRear",
"getFront",
"getFront",
"deleteFront",
"insertFront",
"insertLast",
"insertFront",
"getFront",
"insertFront"]
[[4],[9],[],[],[],[],[],[6],[5],[9],[],[6]]
*/
