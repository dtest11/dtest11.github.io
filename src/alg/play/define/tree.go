package define

// BinaryTreeNode BinaryTreeNode
type BinaryTreeNode struct {
	Data  string
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// TrieNode TrieNode
type TrieNode struct {
	Data string
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
