package stack

func longestValidParentheses(s string) int {
	var st = NewStack()
	st.Push(-1)
	maxMan := 0
	for i := 0; i < len(s); i++ {
		ele := string(s[i])
		if ele == "(" { // 始终将左括号入栈（
			st.Push(i)
		} else {
			st.Pop() // 将栈顶的元素弹出
			if st.Len() == 0 {
				st.Push(i)
			} else {
				tmp := i - st.Peek().(int)
				if tmp > maxMan {
					maxMan = tmp
				}
			}

		}
	}
	return maxMan
}
