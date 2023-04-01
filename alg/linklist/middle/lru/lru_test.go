package lru

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	l := Constructor(10)
	for i := 0; i < 12; i++ {
		l.Put(i, i)
	}
	l.link.Display()
}
