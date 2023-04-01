package string

import (
	"strings"
)

func isValid(s string) bool {
	symbols := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	var stack Stack // 先进先出 FILO
	for _, v := range s {
		v := string(v)
		target, ok := symbols[v]
		if !ok {
			stack.Push(v)
			continue
		}
		if stack.Len() == 0 {
			return false
		}
		top := stack.Pop()
		if top != target {
			return false
		}
	}
	return stack.Len() == 0
}

type Stack struct {
	data []string
}

func (s *Stack) Push(v string) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() string {
	ele := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return ele
}

func (s *Stack) Len() int {
	return len(s.data)
}

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {

		if !isValidStr(s[l]) {
			l++
			continue
		}
		if !isValidStr(s[r]) {
			r--
			continue
		}
		if strings.EqualFold(string(s[l]), string(s[r])) {
			l++
			r--
			continue
		} else {
			return false
		}

	}
	return true
}

func isValidStr(s uint8) bool {
	if s >= 'a' && s <= 'z' {
		return true
	}
	if s >= 'A' && s <= 'Z' {
		return true
	}
	if s >= '0' && s <= '9' {
		return true
	}
	return false
}
