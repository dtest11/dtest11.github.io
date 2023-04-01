reverse list

```go
func reverse(head *ListNode) *ListNode {
	var current = head
	var prev *ListNode

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
```

middle
```go
func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}


```
#### 翻转left-right 中间的数据
定义3个listNode,
l -[left-right]-r
遍历head ，将节点添加到3个节点，然后合并

```go
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	idx := 1
	var h = head
	var m = &ListNode{Val: -1}
	var mRoot = m
	var l = &ListNode{Val: -1} // 存储前半部分
	var r = &ListNode{Val: -1} // 存储后半部分
	var result = l
	for head != nil {
		if idx >= left && idx <= right {
			m.Next = head // 满足条件的数据
			m = m.Next
		} else if idx < left {
			l.Next = head //前半部分数据
			l = l.Next
		} else {
			r.Next = head // 后半部分
			break
		}
		idx++
		head = head.Next
	}
	l.Next = nil
	m.Next = nil
	if mRoot.Next == nil {
		return h
	}
	mRoot = reverse(mRoot.Next)
	l.Next = mRoot
	for mRoot.Next != nil {
		mRoot = mRoot.Next
	}
	mRoot.Next = r.Next
	return result.Next
}

func reverse(node *ListNode) *ListNode {
	cur := node
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

```

. 2个链表的值相加 
```text
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
```
做法：先将2个字符reverse,然后循环遍历累加，对于长度更长的链表，单独处理
这里需要主要的是

val = (a+b+carry) % 10

carry = (a+b+carry) /10
```go

/*
*
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
*/


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	var result = dummy
	carry := 0
	for l1 != nil && l2 != nil {
		val := l2.Val + l1.Val + carry
		dummy.Next = &ListNode{Val: val % 10}
		dummy = dummy.Next
		carry = val / 10
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		val := l1.Val + carry
		dummy.Next = &ListNode{Val: val % 10}
		dummy = dummy.Next
		l1 = l1.Next
		carry = val / 10
	}
	for l2 != nil {
		val := l2.Val + carry
		dummy.Next = &ListNode{Val: val % 10}
		dummy = dummy.Next
		l2 = l2.Next
		carry = val / 10
	}
	if carry != 0 {
		dummy.Next = &ListNode{Val: carry}
		dummy = dummy.Next
	}
	return result.Next
}

```
* 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

方法一：闭合为环
思路及算法
记给定链表的长度为 n，注意到当向右移动的次数 k≥n 时，我们仅需要向右移动 kmodn 次即可。因为每 n 次移动都会让链表变为原状。这样我们可以知道，新链表的最后一个节点为原链表的第 (n−1)−(kmodn) 个节点（从 0 开始计数）。
这样，我们可以先将给定的链表连接成环，然后将指定位置断开。
具体代码中，我们首先计算出链表的长度 n，并找到该链表的末尾节点，将其与头节点相连。这样就得到了闭合为环的链表。然后我们找到新链表的最后一个节点（即原链表的第 (n−1)−(kmodn) 个节点），将当前闭合为环的链表断开，即可得到我们所需要的结果。
特别地，当链表长度不大于 1，或者 k 为 n 的倍数时，新链表将与原链表相同，我们无需进行任何处理。
```go
func rotateRight(head *ListNode, k int) *ListNode {
    if k == 0 || head == nil || head.Next == nil {
        return head
    }
    n := 1
    iter := head
    for iter.Next != nil {
        iter = iter.Next
        n++
    }
    add := n - k%n
    if add == n {
        return head
    }
    iter.Next = head
    for add > 0 {
        iter = iter.Next
        add--
    }
    ret := iter.Next
    iter.Next = nil
    return ret
}


```