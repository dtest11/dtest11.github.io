#BFS 广度搜素

使用队列 来存储每一层的信息
```text
func bfs(node tree) {
    var q Queue
    q.push(node)
    for q !=nil {
       l := q.len()
       for i :=0;i < l;i++ {
           cur :=q.pop()
           fmt.Println(cur.Val)
           q.push(cur.left)
           q.push(cur.right)
       }
    }

}
```


1. [二叉树的最小深度](/bfs/mini_height.go)
