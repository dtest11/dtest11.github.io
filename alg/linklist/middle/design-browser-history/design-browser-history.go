package designbrowserhistory

type Node struct {
	Val  string
	Next *Node
	Prev *Node
}

type BrowserHistory struct {
	history *Node
	cur     *Node
}

func Constructor(homepage string) BrowserHistory {
	head := &Node{Val: homepage}
	return BrowserHistory{
		history: head,
	}
}

func (h *BrowserHistory) Visit(url string) {
	n := &Node{Val: url}
	h.history.Next = n
	n.Prev = h.history
	h.history = h.history.Next
}

func (h *BrowserHistory) Back(steps int) string {
	for h.history != nil && steps != 0 {
		if h.history.Prev == nil {
			break
		}
		h.history = h.history.Prev
		steps--
	}
	return h.history.Val
}

func (h *BrowserHistory) Forward(steps int) string {
	for h.history != nil && steps != 0 {
		if h.history.Next == nil {
			break
		}
		h.history = h.history.Next
		steps--
	}
	return h.history.Val
}
