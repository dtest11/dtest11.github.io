### [171. Excel 表列序号](https://leetcode.cn/problems/excel-sheet-column-number/)

给你一个字符串columnTitle ，表示 Excel 表格中的列名称。返回 该列名称对应的列序号。

例如：
```text


A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...


示例 1:

输入: columnTitle = "A"
输出: 1
示例2:

输入: columnTitle = "AB"
输出: 28
示例3:

输入: columnTitle = "ZY"
输出: 701


```


```go
import "math"

func titleToNumber(n string) int {
	// 26 进制
	var symbols = make(map[rune]int)
	j := 1
	for i := 'A'; i <= 'Z'; i++ {
		symbols[i] = j
		j++
	}
	result := 0
	count := 0

	for i := len(n) - 1; i >= 0; i-- {

		pow := math.Pow(float64(26), float64(count))
		result += int(pow) * (symbols[rune(n[i])])
		count++
	}
	return result
}

```