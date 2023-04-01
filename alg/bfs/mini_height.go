package main

type Tree struct {
	l   *Tree
	r   *Tree
	val int
}

func miniHeight(node *Tree) int {
	if node == nil {
		return 0
	}
	var q []*Tree
	q = append(q, node)
	var level int
	for len(q) != 0 {
		l := len(q)
		level++
		for i := 0; i < l; i++ {
			cur := q[i]
			if cur.l == nil && cur.r == nil { // 没有左右子树
				return level
			}

			if cur.l != nil {
				q = append(q, cur.l)
			}
			if cur.r != nil {
				q = append(q, cur.r)
			}
		}
		// 从q 中删除前l 个元素
		q = append(q[:l], q[l+1:]...)
	}
	return level
}

type hello struct {
	Name string
	Age  string
}

type MyHello struct {
	Name string `json:"name,omitempty"`
	Age  string `json:"age,omitempty"`
}
