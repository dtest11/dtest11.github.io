package geekbang

import (
	"fmt"

	"github.com/9999-dollor/algorithm/define"
)

// https://juejin.im/post/5aa299c1518825557b4c5806

// AddNode 创建一个单向链表
func AddNode(linkList *define.LinkNode, val int) *define.LinkNode {
	temp := define.LinkNode{Val: val}
	if linkList == nil {
		return &temp
	}

	for linkList.Next != nil {
		linkList = linkList.Next
	}
	linkList.Next = &temp
	return linkList
}

// PrintlnLinkList 打印
func PrintlnLinkList(linkList *define.LinkNode) {
	fmt.Printf("翻转单链表: \t")
	for linkList != nil {
		fmt.Printf(" %d  > \t", linkList.Val)
		linkList = linkList.Next
	}
	fmt.Println("")
}

// ReverLinkList 翻转单链表， 先读取出来，然后重新构建
func ReverLinkList(head *define.LinkNode) {
	var datas []int
	for head != nil {
		datas = append(datas, head.Val)
		head = head.Next
	}
	var result *define.LinkNode
	index := len(datas) - 1
	for i := index; i >= 0; i-- {
		if index == i {
			result = AddNode(nil, datas[i])
		} else {
			AddNode(result, datas[i])
		}
	}
	PrintlnLinkList(result)
}

// ReverLinkList2 翻转单链表， 原地转换
func ReverLinkList2(head *define.LinkNode) {
	if head == nil {
		return
	}
	if head.Next == nil {
		return
	}

	var (
		prev *define.LinkNode
		next *define.LinkNode
	)

	// 利用2个指针来翻转
	for head != nil {
		next = head.Next // 记录下一个需要操作的元素
		head.Next = prev // 添加反转的元素到后面
		prev = head      // 缓存
		head = next      //
	}
	PrintlnLinkList(prev)
}

// IsCircle fast 每次走2个节点, slow : 走一个节点, 相遇的时候就是有
func IsCircle(head *define.LinkNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return false
	}
	var (
		slow *define.LinkNode
		fast *define.LinkNode
	)
	slow = head
	fast = head.Next
	for slow != nil && fast != nil {
		if fast.Next == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false

}

// IsCircle2 循环遍历，查找相同的点
func IsCircle2(head *define.LinkNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	current := head.Next
	for current != nil {
		if current == head {
			return true
		}
		current = current.Next
	}
	return false
}

// GetMiddleElement 获取链表的中间元素
// 使用2个快慢指针，快的每次跑2个，慢的跑1个，当快的走到末尾的时候，那么慢的肯定到中间
func GetMiddleElement(head *define.LinkNode) *define.LinkNode {
	if head == nil {
		return nil
	}
	panic("not op")
}

// // MergeLinkList 2个有序链表的合并
// func (firt *LinkNode) MergeLinkList(second *LinkNode) *LinkNode {

// 	for firt != nil || second != nil {
// 		f1 := firt
// 		f2 := f1.Next

// 		s1 := second
// 		s2 := s1.Next

// 		if f1.Val > s1.Val {
// 			s1.Next = f1
// 			second = s2
// 			firt = f2
// 		}

// 		if f1.Val <= s1.Val {
// 			f1.Next = s1
// 			s1.Next = f2
// 			firt = f2
// 			second = s2
// 		}
// 	}
// }
