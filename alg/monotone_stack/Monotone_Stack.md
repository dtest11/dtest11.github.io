单调栈
```go
type st struct {
	
}
func (s *st) Push(ele []int)  {

}
// 查看
func (s *st) Peek()  {

}

// 获取&delete
func (s *st) Pop()  {

}

var current int
var stack st
// 保证单调的性质
for stack !=nil && stack.peek()> current {
	stack.pop()
}
stack.Push(ele)

```