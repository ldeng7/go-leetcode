func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	n := root
	if p.Val > q.Val {
		p, q = q, p
	}
	for {
		if p.Val <= n.Val && q.Val >= n.Val {
			return n
		}
		if q.Val <= n.Val {
			n = n.Left
		} else {
			n = n.Right
		}
	}
	return nil
}
