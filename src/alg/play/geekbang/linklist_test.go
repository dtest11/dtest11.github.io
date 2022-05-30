package geekbang

import (
	"fmt"
	"testing"

	"github.com/9999-dollor/algorithm/define"
)

func TestPrintlnLinkList(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5}
	var linkNode define.LinkNode

	for _, v := range vals {
		fmt.Println(v)
		AddNode(&linkNode, v)
	}
	PrintlnLinkList(&linkNode)
	// 反转
	ReverLinkList(&linkNode)

	ReverLinkList2(&linkNode)
}

func TestIsCircle(t *testing.T) {
	vals := []int{2, 3, 4, 5}
	var linkNode = &define.LinkNode{Val: 1, Next: nil}

	for _, v := range vals {
		AddNode(linkNode, v)
	}
	var head = linkNode

	for linkNode.Next != nil {
		linkNode = linkNode.Next
	}
	linkNode.Next = head // 循环列表
	result := IsCircle(linkNode)
	//result := IsCircle2(linkNode)
	fmt.Println(result)
}
