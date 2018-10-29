func check(s, t *TreeNode) bool {
	if nil == s && nil == t {
		return true
	} else if nil == s || nil == t {
		return false
	} else if s.Val != t.Val {
		return false
	}
	return check(s.Left, t.Left) && check(s.Right, t.Right)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if nil == s {
		return false
	} else if check(s, t) {
		return true
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}
