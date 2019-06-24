func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	var cal func(*TreeNode, int) bool
	cal = func(n *TreeNode, s int) bool {
		if nil == n.Left && nil == n.Right {
			return s+n.Val < limit
		}
		bl, br := true, true
		if nil != n.Left {
			bl = cal(n.Left, s+n.Val)
		}
		if nil != n.Right {
			br = cal(n.Right, s+n.Val)
		}
		if bl {
			n.Left = nil
		}
		if br {
			n.Right = nil
		}
		return bl && br
	}
	if cal(root, 0) {
		return nil
	}
	return root
}
