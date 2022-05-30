package stack

import (
	"fmt"
	"testing"
)

func isValid(s string) bool {
	preDefine := map[string]string{")": "(", "]": "[", "}": "{"}
	var result = NewStack()
	for _, v := range s {
		c := string(v)
		if v1, ok := preDefine[c]; !ok {
			result.Push(c)
		} else if result.Len() != 0 {
			top := result.Peek().(string)
			if top != v1 {
				return false
			}
			result.Pop()
		} else {
			return false
		}
	}

	return result.Len() == 0
}

func Test_vald(t *testing.T) {
	data := "()[]"
	fmt.Println(isValid(data))
	data = "()[]{}"
	fmt.Println(isValid(data))

	data = "({[)"
	fmt.Println(isValid(data))
}
