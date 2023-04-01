package graph

import "math"

/**
最短路径算法
连通地图上的点
**/

type Node struct {
	Id            int
	distFromStart int
}

func dijkstra(start int, graph [][3]int) []int {
	distTo := make([]int, len(graph)+1)
	for i := range distTo {
		distTo[i] = math.MaxInt
	}
	distTo[start] = 0
	var pq *PriorityQueue
	for pq.len() != 0 {
		cur := pq.Pop()
		id := cur.Id
		if cur.distFromStart < distTo[id] {
			continue // 当前值已经是最小的了，跳过本次
		}
		// 找到id的兄弟节点
		for _, neighbor := range graph[id] {
			nextNodeId := neighbor[0]
			distToNextId := neighbor[1] + distTo[id]
			if distToNextId < distTo[nextNodeId] {
				distTo[nextNodeId] = distToNextId
				pq.Push(Node{Id: nextNodeId, distFromStart: distToNextId})
			}
		}
	}
	return distTo
}

type PriorityQueue struct{}

func (p *PriorityQueue) len() int {
	panic("not implement")
}

func (p *PriorityQueue) Pop() Node {
	panic("not implement")
}

func netWorkDelayTimes(times [][]int, n int, k int) int {
	var graph [][3]int = make([][3]int, k)
	for index, v := range times {
		graph[index] = append(v[0], v[1], v[2])
	}
	var distTo = dijkstra(k, graph)
	max := 0
	for _, v := range distTo {
		if v == math.MaxInt {
			return -1
		}
		if v > max {
			max = v
		}
	}
	return max
}
func (p *PriorityQueue) Push(node Node) {
	panic("not implement")
}
