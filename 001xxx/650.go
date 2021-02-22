func lowestCommonAncestor(p *Node, q *Node) *Node {
	m := map[int]bool{}
	for nil != p {
		m[p.Val] = true
		p = p.Parent
	}
	for nil != q {
		if m[q.Val] {
			return q
		}
		q = q.Parent
	}
	return nil
}
