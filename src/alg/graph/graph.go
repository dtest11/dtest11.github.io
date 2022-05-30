package graph

/**
graph: 就是邻接表， 寻找停止的条件，所有可能的路径， 暴力搜素
然后路径的长度 已经是所有的元素的时候停止就行
*/
func allPathsSourceTarget(graph [][]int) [][]int {
	var result [][]int
	allPathsSourceTargetHelp(graph, 0, len(graph), nil, &result)
	return result
}

func allPathsSourceTargetHelp(graph [][]int, index int, length int, tmp []int, result *[][]int) {
	if index == length {
		*result = append(*result, tmp)
		// tmp remove last
	}
	for _, v := range graph[index] {
		tmp = append(tmp, v)
		allPathsSourceTargetHelp(graph, index+1, length, tmp, result)
	}
}

type Tree struct {
	val      int
	children []*Tree
}

//  多叉树
func muiltTreeDFS(root *Tree) {
	if root == nil {
		return
	}
	for _, v := range root.children {
		muiltTreeDFS(v)
	}
}
