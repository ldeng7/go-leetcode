func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
	if nil == root1 {
		return false
	}
	var cal func(*TreeNode, int) bool
	cal = func(n *TreeNode, v int) bool {
		if nil == n {
			return false
		} else if n.Val == v {
			return true
		} else if n.Val < v {
			return cal(n.Right, v)
		}
		return cal(n.Left, v)
	}
	return cal(root2, target-root1.Val) || twoSumBSTs(root1.Left, root2, target) ||
		twoSumBSTs(root1.Right, root2, target)
}
