package add_two_numbers

import (
	"fmt"
	"testing"

	"github.com/dtest11/alg/linklist/build"
)

type Case struct {
	inputA []int
	inputB []int
	expect []int
}

func Test_addTwoNumbers(t *testing.T) {
	case1 := Case{
		inputA: []int{2, 4, 9},
		inputB: []int{5, 6, 4, 9},
		expect: []int{7, 0, 4, 0, 1},
	}

	newExample(case1, t)

	case1 = Case{
		inputA: []int{9, 9, 9, 9, 9, 9, 9},
		inputB: []int{9, 9, 9, 9},
		expect: []int{7, 0, 4, 0, 1},
	}

	newExample(case1, t)
}

func newExample(params Case, t *testing.T) {
	la := build.NewLinKList()

	la.InsertFronts(params.inputA)
	lb := build.NewLinKList()
	lb.InsertFronts(params.inputA)
	result := addTwoNumbers(la.Head.Next, lb.Head.Next)
	var temp []int
	for result != nil {
		temp = append(temp, result.Val)
		result = result.Next
	}
	t.Log("**********")
	t.Logf("result:%v", temp)
	t.Logf("expect:%v", params.expect)
	t.Log("**********")
}

func Test_deleteMiddle(t *testing.T) {
	l := build.NewLinKList()
	//l.InsertFronts([]int{1, 3, 4, 7, 1, 2, 6})
	l.InsertFronts([]int{2, 1})
	res := deleteMiddle(l.Head.Next)
	fmt.Println(res)
}

func Test_addTwoNumbersV3(t *testing.T) {
	l1 := build.NewLinKList()
	l1.InsertFronts([]int{7, 2, 4, 3})

	l2 := build.NewLinKList()
	l2.InsertFronts([]int{5, 6, 4})

	addTwoNumbersV3(l1.Head.Next, l2.Head.Next)

}

func Test_numComponents(t *testing.T) {
	l2 := build.NewLinKList()
	l2.InsertFronts([]int{0, 1, 2, 3})

	fmt.Println(numComponents(l2.Head.Next, []int{0, 1, 3}))
}
