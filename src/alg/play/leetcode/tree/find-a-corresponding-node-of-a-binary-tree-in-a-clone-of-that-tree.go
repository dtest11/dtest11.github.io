package tree

//medium

func getTargetCopy(original, cloned *TreeNode) *TreeNode {
	if original == nil {
		return nil
	}
	if original == cloned {
		return original
	}

	if result := getTargetCopy(original.Left, cloned.Left); result != nil {
		return result
	}
	return getTargetCopy(original.Right, cloned.Right)
}
