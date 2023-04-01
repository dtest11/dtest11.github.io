package spiralmatrixiv

import (
	"fmt"
	"github.com/dtest11/alg/linklist/build"
	swappingnodesinalinkedlist "github.com/dtest11/alg/linklist/middle/swapping-nodes-in-a-linked-list"
	"testing"
)

func Test_spiralMatrix(t *testing.T) {
	//data := []int{3, 0, 2, 6, 8, 1, 7, 9, 4, 2, 5, 5, 0}
	data := []int{1, 2, 3, 4, 5}
	link := build.NewLinKList()
	link.InsertFronts(data)
	res := spiralMatrix(3, 5, link.Head.Next)
	fmt.Println(res)
	swappingnodesinalinkedlist.SwapNodes(link.Head.Next, 2)

}
