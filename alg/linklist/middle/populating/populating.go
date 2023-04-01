package populating

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 层次遍历然后连接
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	for len(queue) != 0 {

		tmp := queue
		queue = nil
		for i := 0; i < len(tmp); i++ {
			if i+1 <= len(tmp)-1 {
				tmp[i].Next = tmp[i+1]
			}

			cur := tmp[i]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

	}
	return root
}
