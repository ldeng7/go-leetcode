func cal(n *TreeNode, val int, cnt *int) bool {
	if nil == n {
		return true
	}
	bl, br := cal(n.Left, n.Val, cnt), cal(n.Right, n.Val, cnt)
	if !(bl && br) {
		return false
	}
	(*cnt)++
	return n.Val == val
}

func countUnivalSubtrees(root *TreeNode) int {
	cnt := 0
	cal(root, -1, &cnt)
	return cnt
}
