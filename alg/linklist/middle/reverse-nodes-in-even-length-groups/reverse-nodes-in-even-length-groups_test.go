package reversenodesinevenlengthgroups

import (
	"fmt"
	"testing"

	"github.com/dtest11/alg/linklist/build"
)

func TestReverr(t *testing.T) {
	l := build.NewLinKList()
	l.InsertFronts([]int{5, 2, 6, 3, 9, 1, 7, 3, 8, 4})
	res := reverseEvenLengthGroups(l.Head.Next)
	fmt.Println(res)
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}

	l.ClearWith([]int{1, 1, 0, 6, 5})
	res = reverseEvenLengthGroups(l.Head.Next)
	fmt.Println(res)
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}

	l.ClearWith([]int{4, 3, 0, 5, 1, 2, 7, 8, 6})
	res = reverseEvenLengthGroups(l.Head.Next)
	fmt.Println(res)
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
}
