package geekbang

import (
	"testing"
)

func Test_newlru(t *testing.T) {
	l := newlru(10)

	l.set(1, 1)
	a := l.get(1)
	t.Log(a)
}
