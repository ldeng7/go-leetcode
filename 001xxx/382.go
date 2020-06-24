func balanceBST(root *TreeNode) *TreeNode {
	ar := make([]*TreeNode, 0, 32)
	var build func(int, int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		m := (l + r) >> 1
		n := ar[m]
		n.Left, n.Right = build(l, m-1), build(m+1, r)
		return n
	}
	var cal func(*TreeNode)
	cal = func(n *TreeNode) {
		if nil == n {
			return
		}
		cal(n.Left)
		ar = append(ar, n)
		cal(n.Right)
	}
	cal(root)
	return build(0, len(ar)-1)
}
