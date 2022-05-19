# BFS

word ladder: [https://www.youtube.com/watch?v=hB\_nYXFtwP0](https://www.youtube.com/watch?v=hB\_nYXFtwP0)

拓扑排序：

[https://www.youtube.com/watch?v=B5hxqxBL2d0](https://www.youtube.com/watch?v=B5hxqxBL2d0)

# $DFS

```go

func Dfs(root Node) {
    if root == nil {
        return 
    }
    if root visited {
        return 
    }
    for child in root {
        if child visted {
            continue
        }
        Dfs(child)
    }
}

```

# $ Bfs 

```go

func Bfs(root Node){
    if root == nil {
        return 
    }
    var queue Queue
    queue.Push(root)
    for !queue.IsEmpty(){
       size :=queue.Size()
       for i :=0; i<size; i++{
           cur :=queue.Pop()
           if cur.Left !=nil {
               queue.Push(cur.Left)
           }
           if cur.Right !=nil {
               queue.Push(cur.Right)
           }
       }
    }
}




```