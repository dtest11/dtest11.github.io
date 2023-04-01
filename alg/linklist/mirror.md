## [1367. 二叉树中的链表](https://leetcode.cn/problems/linked-list-in-binary-tree/)

给你一棵以 root 为根的二叉树和一个 head 为第一个节点的链表。

如果在二叉树中，存在一条一直向下的路径，且每个点的数值恰好一一对应以 head 为首的链表中每个节点的值，那么请你返回 True ，否则返回 False 。

一直向下的路径的意思是：从树中某个节点开始，一直连续向下的路径。

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/02/29/sample_1_1720.png)

```text
输入：head = [4,2,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
输出：true
解释：树中蓝色的节点构成了与链表对应的子路径。

```

```go
func isSubPath(head *ListNode, root *TreeNode) bool {
   if head == nil {
       return true
   }
   if root == nil {
       return false
   }
   return isSub(head,root) || isSubPath(head,root.Left) || isSubPath(head,root.Right)
}

func isSub(head *ListNode, root *TreeNode) bool {
    if head == nil {
        return true
    }
    if root ==nil {
        return false
    }
    if head.Val != root.Val {
        return false
    }
    return isSub(head.Next,root.Left) || isSub(head.Next,root.Right)
}
```