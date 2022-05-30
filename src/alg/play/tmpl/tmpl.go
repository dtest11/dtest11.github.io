package stack

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

type Stack struct {
	top    []interface{}
	length int
}

// Create a new stack
func NewStack() *Stack {
	return &Stack{nil, 0}
}

// Return the number of items in the stack
func (st *Stack) Len() int {
	return st.length
}

// Return the number of items in the stack
func (st *Stack) Data() interface{} {
	return st.top
}

func (st *Stack) Empty() bool {
	return st.length == 0
}

// View the top item on the stack
func (st *Stack) Peek() interface{} {
	return st.top[st.length-1]
}

// Pop the top item of the stack and return it
func (st *Stack) Pop() interface{} {
	index := st.Len() - 1    // Get the index of the top most element.
	element := st.top[index] // Index into the slice and obtain the element.
	st.top = st.top[:index]
	st.length--
	return element
}

// Push a value onto the top of the stack
func (st *Stack) Push(value interface{}) {
	st.top = append(st.top, value)
	st.length++
}

//逆波兰表达式计算： 遇到数子，进栈， 遇到非数字，弹出栈顶的2个元素参与计算
// 什么是逆波兰表达式 1+2 => 1，2 *
func evalRPN(tokens []string) int {
	st := NewStack()
	for _, v := range tokens {
		switch v {
		case "+":
			n1 := st.Pop().(int)
			n2 := st.Pop().(int)
			st.Push(n2 + n1)
		case "-":
			n1 := st.Pop().(int)
			n2 := st.Pop().(int)
			st.Push(n2 - n1)
		case "*":
			n1 := st.Pop().(int)
			n2 := st.Pop().(int)
			st.Push(n2 * n1)
		case "/":
			n1 := st.Pop().(int)
			n2 := st.Pop().(int)
			st.Push(n2 / n1)
		default:
			// must number
			n, ok := strconv.ParseInt(v, 10, 64)
			if ok != nil {
				panic(ok)
			}
			st.Push(int(n))
		}
	}
	if st.Len() != 0 {
		return st.Pop().(int)
	}
	return 0
}

func getPolish(s string) []string {
	var result []string
	var operators = NewStack()
	var priors = map[string]int{"+": 1, "-": 1, "*": 2, "/": 2, "^": 3}
	for i := 0; i < len(s); {
		v := string(s[i])
		if v == " " {
			continue
		}
		log.Println("结果集：", result)
		log.Println("操作符集合：", operators.Data())
		switch v {
		case "+", "-", "*", "/", "^":
			log.Println("operator:", v)
			for !operators.Empty() {
				p1 := priors[v]
				top := priors[operators.Peek().(string)]
				// 棧頂的元素是否是 更高優先級
				if top >= p1 {
					result = append(result, operators.Pop().(string))
					//result.Push(operators.Pop())
					continue
				} else {
					break
				}
			}

			operators.Push(v)
			log.Printf("operator 推入棧頂:%v \n", operators.Data())
			i += 1

		case "(":
			log.Println(`(左括号，也是直接入栈`)
			operators.Push(v)
			i += 1

		case ")":
			log.Println(`)右括号的话,从运算符栈中弹出数据,直到遇到左括号(`)
			for !operators.Empty() {
				if operators.Peek().(string) == "(" {
					operators.Pop()
					break
				} else {
					//result.Push(operators.Pop())
					result = append(result, operators.Pop().(string))

				}
			}
			i += 1

		default:
			//入队的时候要判断后面是否还有剩余的数字，要把整个数字入队列，而不是一个数字字符
			var tmp string
			for j := i; j < len(s); j++ {
				if isDigit(string(s[j])) {
					tmp += string(s[j])
				}
			}

			result = append(result, tmp)

			log.Printf("操作数 直接压入栈:%v \n", result)
			i += len(tmp)
		}
	}
	for !operators.Empty() {
		//result.Push(operators.Pop())
		result = append(result, operators.Pop().(string))

	}
	log.Println("result", result)
	return result

}

func isDigit(v string) bool {
	n := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, vv := range n {
		if v == vv {
			return true
		}
	}
	return false
}

func Test_evaluate(t *testing.T) {
	s := "356+4*2/(1-5)^2^3"
	//s="(1+(4+5+2)-3)+(6+8)"
	// s := "3+4*2/(1-5)"
	fmt.Println(getPolish(s))
}

///***
//我们从左到右遍历表达式，根据遍历到的元素分成下面几种情况：
//1、如果是操作数，那么直接压入栈
//2、如果是 '(" 左括号的话，也是直接入栈
//3、如果是“）”右括号的话，那么就从运算符栈中弹出数据，并入队列，直到遇见“（”左括号 （注意，这里的左括号只是从栈中弹出，但是并不入队列），右括号“）”不入栈。
//4、如果是算术运算符，也就是“+-*/”这些，就从栈顶开始，依次弹出比当前处理的运算符优先级高和优先级相等的运算符入队列，
//3直到一个比他优先级低的或遇到了一个“（”左括号为止，
//然后将当前正在处理的运算符入栈。
//5、重复上述操作，直到遍历到中缀表达式的最后一个符号。此时队列中存放的便是后缀表达式。
// */

// 先将数据 转化成逆波兰表达式， 然后计算
func calculate(s string) int {
	rpn := getPolish(s)
	result := evalRPN(rpn)
	return result
}