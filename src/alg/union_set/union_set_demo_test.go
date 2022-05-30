package union_set

import (
	"fmt"
	"testing"
)

func Test_equationsPossible(t *testing.T) {
	var input = []string{"a==b", "b!=c", "c==a"}
	fmt.Println(equationsPossible(input))

	input = []string{"c==c", "b==d", "x!=z"}
	fmt.Println(equationsPossible(input))
}

/**
go test github.com/dtest11/dump.sort/union_set  -timeout 30s -v -run ^Test_equationsPossible$
**/
