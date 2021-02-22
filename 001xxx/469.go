func getLonelyNodes(root *TreeNode) []int {
	o := make([]int, 0, 16)
	var cal func(*TreeNode)
	cal = func(n *TreeNode) {
		if nil == n {
			return
		}
		l, r := n.Left, n.Right
		if nil != l && nil == r {
			o = append(o, l.Val)
		}
		if nil != r && nil == l {
			o = append(o, r.Val)
		}
		cal(l)
		cal(r)
	}
	cal(root)
	return o
}
