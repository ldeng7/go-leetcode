func countUnivalSubtrees(root *TreeNode) int {
	cnt := 0
	var cal func(*TreeNode, int) bool
	cal = func(n *TreeNode, val int) bool {
		if nil == n {
			return true
		}
		bl, br := cal(n.Left, n.Val), cal(n.Right, n.Val)
		if !(bl && br) {
			return false
		}
		cnt++
		return n.Val == val
	}
	cal(root, -1)
	return cnt
}
