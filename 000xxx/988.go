func smallestFromLeaf(root *TreeNode) string {
	var cal func(n *TreeNode, s string) string
	cal = func(n *TreeNode, s string) string {
		if nil == n {
			return ""
		}
		s = string('a'+n.Val) + s
		l, r := cal(n.Left, s), cal(n.Right, s)
		ll, lr := len(l), len(r)
		if ll+lr == 0 {
			return s
		} else if ll*lr == 0 {
			if 0 == ll {
				return r
			}
			return l
		} else if l < r {
			return l
		}
		return r
	}
	return cal(root, "")
}
