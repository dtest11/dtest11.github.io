### 图的表示

图的表示方法

1. 邻接表

```text
1 -> 2 ->3 -> 4
2 -> 2 -3->5 -6

```

3. 二维矩阵

```text
1 2 3 4
3 2 1 5 
5 5 6 7
```

优缺点：
邻接表相比二维矩阵 需要存储的空间小
二维矩阵可以快速找到相邻的节点

## 2， 图的遍历

邻接矩阵

```go
func travle(num []LinkList, index, travelRecord []Node) {
node := num[index]
if travelRecord[node] {
return // 已经访问过
}
for i  range node{
travel(num, i, travelRecord),
}
}


func build(input [][2]int) [][]int{
	var result [][]int
	for _,v :=range input{
		index :=v[0] // 可能不是连续，遍历的时候需要考虑
		val :=v[1]
		result[index]= append(result[index], val)
	}
	return result
}
```


