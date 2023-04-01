package v1

import (
	v2 "github.com/dtest11/alg/lfu_lru/lfu/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	//l := Constructor(2)
	l := v2.Constructor(2)
	//l1 := `["LFUCache","put","put","get","put","get","get","put","get","get","get"]`
	//l2 := `[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]`
	//l3 := `[null,null,null,1,null,-1,3,null,-1,3,4]`
	//list := strings.Split(l1, ",")
	//for i := 0; i < len(list); i++ {
	//	fmt.Println(list[i])
	//	fmt.Println(strings.Split(l2, ",[")[i])
	//	fmt.Println(strings.Split(l3, ",")[i])
	//
	//	fmt.Println("--------------------------------------------")
	//}
	l.Put(1, 1)
	l.Put(2, 2)
	assert.Equal(t, 1, l.Get(1))
	l.Put(3, 3)
	assert.Equal(t, -1, l.Get(2))
	assert.Equal(t, 3, l.Get(3))
	l.Put(4, 4)
	assert.Equal(t, -1, l.Get(1))
	assert.Equal(t, 3, l.Get(3))
	assert.Equal(t, 4, l.Get(4))
}
