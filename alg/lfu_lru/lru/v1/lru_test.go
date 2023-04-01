package v1

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	l := Constructor(10)
	for i := 0; i < 12; i++ {
		l.Put(i, i)
	}
	for e := l.list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func Test_newBestLinkList(t *testing.T) {
	b := newBestLinkList()
	for i := 1; i < 10; i++ {
		b.pushFront(&Node{
			Key:   i,
			Value: i,
		})
	}
	b.Display()
	b.removeTail()
	b.Display()

}
