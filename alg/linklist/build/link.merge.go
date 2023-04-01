package build

import (
	"golang.org/x/exp/constraints"
)

type link[T constraints.Ordered] struct {
	Next *link[T]
	data T
}

func NewNode[v constraints.Ordered](data v) *link[v] {
	return &link[v]{data: data}
}

func mergeTwoLink[T constraints.Ordered](a *link[T], b *link[T], dup T) *link[T] {
	curA := a
	curB := b
	var dummy = NewNode(dup)
	for curB != nil && curA != nil {
		eleA := curA.data
		eleB := curB.data
		if eleA >= eleB {
			dummy.Next = NewNode(eleB)
			curB = curB.Next
		} else {
			dummy.Next = NewNode(eleA)
			curA = curA.Next
		}
	}

	if curB != nil {
		dummy.Next.Next = curB // todo nil
	}
	if curA != nil {
		dummy.Next.Next = curA // todo nil
	}
	return dummy.Next
}

// TODO
func mergeTwoSortedLink[v constraints.Ordered](a *link[v], b *link[v]) {

}

func lastKNode[T constraints.Ordered](l *link[T], k int) *link[T] {
	if l == nil {
		return nil
	}
	// 先遍历到K: k+(n-k)
	copyL := l
	index := 0
	for index < k && l != nil {
		l = l.Next
		index++
	}
	for copyL != nil && l != nil {
		copyL = copyL.Next
		l = l.Next
	}
	return copyL
}

func findMiddle[T constraints.Ordered](l *link[T]) *link[T] {
	low := l
	fast := l

	if l == nil {
		return nil
	}

	for low != nil && fast.Next != nil {
		low = low.Next
		fast = fast.Next.Next
	}
	return low
}

func hasCircle[T constraints.Ordered](l *link[T]) bool {
	low := l
	fast := l

	if l == nil {
		return false
	}

	for low != nil && fast.Next != nil {
		low = low.Next
		fast = fast.Next.Next
		if low == fast {
			return true
		}
	}
	return false
}

func findCircle[T constraints.Ordered](l *link[T]) *link[T] {
	if hasCircle(l) {
		low := l
		fast := l
		head := l

		for low != nil && fast.Next != nil {
			low = low.Next
			fast = fast.Next.Next
			if low == fast {
				break
			}
		}
		fast = l
		for head != fast {
			head = head.Next
			fast = fast.Next
		}
		return fast
	}
	return nil
}

func deleteLastKNode[T constraints.Ordered](l *link[T], k int, dump T) *link[T] {
	// find K
	var dummy = NewNode(dump)
	dummy.Next = l
	x := lastKNode(l, k+1)
	x.Next = x.Next.Next
	return dummy.Next

}
