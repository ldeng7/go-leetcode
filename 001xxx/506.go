func findRoot(tree []*Node) *Node {
	v := 0
	for _, n := range tree {
		v ^= n.Val
		for _, c := range n.Children {
			v ^= c.Val
		}
	}
	for _, n := range tree {
		if v == n.Val {
			return n
		}
	}
	return nil
}
