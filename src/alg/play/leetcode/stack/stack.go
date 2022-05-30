package stack

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
