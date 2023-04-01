### [1171. 从链表中删去总和值为零的连续节点](https://leetcode.cn/problems/remove-zero-sum-consecutive-nodes-from-linked-list/)

这里的技巧是使用前缀和【前缀和：表示数据[i],p[0:i-1],p[i:i-3]等的数组
p[i]=p[i-1]+array[i]
```go
func removeZeroSumSublists(head *ListNode) *ListNode {
	var sumMap = make(map[int]*ListNode)
	var sum = 0
	// 因为头结点也有可能会被消掉，所以这里加一个虚拟节点作为头结点
	var dummy = &ListNode{Val: 0}
	dummy.Next = head
	var cur = dummy
	for dummy != nil {
		sum = dummy.Val + sum
		sumMap[sum] = dummy
		dummy = dummy.Next
	}
	sum = 0
	var result = cur
	for cur != nil {
		sum = sum + cur.Val
		next, ok := sumMap[sum]
		if ok {
			cur.Next = next.Next
		}
		cur = cur.Next
	}
	return result.Next
}
```