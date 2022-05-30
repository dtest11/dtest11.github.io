package union_set

import "fmt"

// UnionSet 并查集
type UnionSet struct {
	parents []int
	size    []int
}

func initUnion(n int) *UnionSet {
	var u UnionSet
	for i := 0; i < n; i++ {
		u.parents = append(u.parents, 0)
	}

	for i := 0; i < n; i++ {
		u.size = append(u.size, 0)
	}
	return &u
}

func (u *UnionSet) find(x int) int {
	for u.parents[x] != x { // 说明parent
		x = u.parents[x] // 一直向上找parent
	}
	return x
}

// connected 判断X,Y 是否连通，在一个集合中
func (u *UnionSet) connected(x, y int) bool {
	xParent := u.find(x)
	yParent := u.find(y)
	if xParent != yParent {
		return false
	}
	return true
}

func (u *UnionSet) union(x, y int) {
	fmt.Println(x, y)
	xParent := u.find(x)
	yParent := u.find(y)
	//if xParent != yParent {
	//	u.parents[xParent] = yParent
	//}
	// 优化层级关系， 根据size
	if u.size[xParent] > u.size[yParent] {
		// 小的接到大的下面
		u.parents[yParent] = xParent
		u.size[yParent] += u.size[xParent]

	} else {
		u.parents[xParent] = yParent
		u.size[xParent] += u.size[yParent]
	}
}
