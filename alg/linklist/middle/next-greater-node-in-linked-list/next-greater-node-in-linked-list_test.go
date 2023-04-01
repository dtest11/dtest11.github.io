package next_greater_node_in_linked_list

import (
	"fmt"
	"github.com/dtest11/alg/linklist/build"
	"testing"
)

func Test_nextLargerNodes(t *testing.T) {
	l := build.NewLinKList()
	l.InsertFronts([]int{2, 1, 5})
	res := nextLargerNodes(l.Head.Next)
	fmt.Println(res)
}
