func check(node *TreeNode) bool {
	if nil == node || 0 != node.Val {
		return false
	}
	if nil == node.Left && nil == node.Right {
		return true
	}
	if nil != node.Left && !check(node.Left) {
		return false
	}
	if nil != node.Right && !check(node.Right) {
		return false
	}
	return true
}

func pruneTree(root *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}
	if check(root.Left) {
		root.Left = nil
	} else {
		pruneTree(root.Left)
	}
	if check(root.Right) {
		root.Right = nil
	} else {
		pruneTree(root.Right)
	}
	return root
}
