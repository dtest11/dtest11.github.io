判断是否有环
```go
// hasCycle 快慢指针，是否相遇，注意判断停止的条件
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// detectCycle 先检查是否有环
// 在发现环的地方停止，然后一个节点从开头开始遍历，每个节点每次走一步，然后相遇的地方就是环的开始节点
/***
  https://leetcode.cn/problems/linked-list-cycle-ii/solution/linked-list-cycle-ii-kuai-man-zhi-zhen-shuang-zhi-/
  环的长度定义为b, 从head 到环的长度为a
  根据：
   
  f=2s （快指针每次2步，路程刚好2倍）

  f = s + nb (相遇时，刚好多走了n圈）

  推出：s = nb

  从head结点走到入环点需要走 ： a + nb， 而slow已经走了nb，那么slow再走a步就是入环点了。

  如何知道slow刚好走了a步？ 从head开始，和slow指针一起走，相遇时刚好就是a步
 */
func detectCycle(head *ListNode) *ListNode {
	if !hasCycle(head) {
		return nil
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	slow = head
	for slow != nil && fast != nil {
		if slow == fast {
			return slow
		}
		slow = slow.Next
		fast = fast.Next
	}
	return nil
}

```