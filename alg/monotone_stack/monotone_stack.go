package monotone_stack

// 给一个字符串可能是重复的，删除其中重复出现的字符，并且保证返回的字典序最小
func removeDumplicate(data string) {
	count := make(map[rune]int)
	for _, v := range data {
		count[v]++
	}
	stack := Stack{}
	install := map[rune]bool{}
	for _, v := range data {
		count[v]--
		if install[v] {
			continue
		}

		for !stack.Empty() && stack.Peek() > v {
			cur := stack.Peek()
			if count[cur] == 0 { // 后面没有这个元素了，不能继续删除
				break
			}
			stack.Pop()
			install[cur] = false
		}
		stack.Push(v)
	}
}

// Stack 先进后出
type Stack struct {
	data []rune
}

func (s *Stack) Push(v rune) {
	s.data = append(s.data, v)
}

func (s *Stack) Peek() rune {
	return s.data[len(s.data)-1]
}

func (s *Stack) Pop() rune {
	ele := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return ele
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}
