package dp

import (
	"testing"
)

// go test -timeout 30s github.com/9999-dollor/algorithm/DP -run ^Test_fib$ -v

func Test_fib(t *testing.T) {
	t.Log(fib(10))
	t.Log(fib2(10))

}
