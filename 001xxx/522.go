func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func diameter(root *Node) int {
	o := 0
	var cal func(*Node) (int, int)
	cal = func(n *Node) (int, int) {
		if nil == n {
			return 0, 0
		}
		a, b := 0, 0
		for _, c := range n.Children {
			a1, b1 := cal(c)
			m := max(a1, b1) + 1
			if m > a {
				a, b = m, a
			} else if m > b {
				b = m
			}
		}
		o = max(o, a+b)
		return a, b
	}

	cal(root)
	return o
}
