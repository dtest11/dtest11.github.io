## defer 和return

return 中返回的值 是不受defer 中的逻辑影响的，具体看个例子, return 是先执行的

```go
package main

import "fmt"

func record() int {
    var counter = 10 // counter = 10

    defer func() {
        counter += 10
        fmt.Println("defer中执行counter=", counter) // 这里的counter 不影响return 返回的
    }()
    return counter // counter = 10
}

func main() {
    res := record()
    fmt.Println("最后返回的结果：", res)
}

/****

defer中执行counter= 20
最后返回的结果： 10

*/
```

## defer 中传递参数规则

### 传递函数立即求值
```go
func cal(a, b int) int {
	fmt.Println(a, b)
	return a + b
}

func main() {
	a, b := 3, 4

	defer cal(cal(a, b), a)
	//里面 cal(a,b)在执行到defer的时候会立即求值，确定参数
	// 后续对a,b的修改不会影响a,b,也不会影响外层cal
	a = 100
	b = 8

	fmt.Println("hello")
}

/**
3 4
hello
7 3

*/
```

### not work 
```go
func f2() (r int) { 
    t := 5 
    defer func() { 
        t = t + 5 
    }()
    return t
}
```

### 函数方式传递参数是 传递指针的拷贝
```go
package main

import  "fmt"

func add(a, b int, note int) {
	fmt.Printf("第%d次调用:%d, %d \n", note, a, b)
}

func main() {
	a, b := 3, 4

	defer add(a, b, 0) // 3,4

	defer func() {
		add(a, b, 1) // 7,8
	}() //传递的是参数指针== 值拷贝：拷贝的参数指针

	defer func(a, b int) {
		add(a, b, 2) // 7,8
	}(a, b) //传递的是参数指针== 值拷贝：拷贝的参数指针

	a = 7
	b = 8

	defer add(a, b, 3)
}
```

