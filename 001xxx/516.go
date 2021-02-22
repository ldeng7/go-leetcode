func nodesRemoveOne(ar []*Node, n *Node) []*Node {
	for i, n1 := range ar {
		if n == n1 {
			if i != len(ar)-1 {
				copy(ar[i:], ar[i+1:])
			}
			return ar[:len(ar)-1]
		}
	}
	return ar
}

func nodesInsert(ar []*Node, i int, n *Node) []*Node {
	ar = append(ar, nil)
	if i != len(ar)-1 {
		copy(ar[i+1:], ar[i:])
	}
	ar[i] = n
	return ar
}

func moveSubTree(root *Node, p *Node, q *Node) *Node {
	var cal func(*Node, *Node) [4]*Node
	cal = func(n, f *Node) [4]*Node {
		ns := [4]*Node{}
		if nil == n {
			return ns
		} else if n.Val == p.Val || n.Val == q.Val {
			if n.Val == p.Val {
				ns[0], ns[2] = p, f
			} else {
				ns[1], ns[3] = q, f
			}
			return ns
		}
		for _, c := range n.Children {
			ns1 := cal(c, n)
			if nil != ns1[0] {
				ns[0], ns[2] = ns1[0], ns1[2]
			}
			if nil != ns1[1] {
				ns[1], ns[3] = ns1[1], ns1[3]
			}
		}
		return ns
	}

	ns := cal(root, nil)
	if nil != ns[0] && nil != ns[1] {
		if ns[2].Val != q.Val {
			ns[2].Children = nodesRemoveOne(ns[2].Children, ns[0])
			ns[1].Children = append(ns[1].Children, ns[0])
		}
		return root
	}

	var cal1 func(*Node, *Node, *Node) [2]*Node
	cal1 = func(n, m, f *Node) [2]*Node {
		ns := [2]*Node{}
		if nil == n {
			return ns
		} else if n.Val == m.Val {
			ns[0], ns[1] = m, f
			return ns
		}
		for _, c := range n.Children {
			ns1 := cal1(c, m, n)
			if nil != ns1[0] {
				ns[0], ns[1] = ns1[0], ns1[1]
			}
		}
		return ns
	}

	if nil != ns[1] {
		ns1 := cal1(ns[1], p, nil)
		if ns1[1].Val != q.Val {
			ns1[1].Children = nodesRemoveOne(ns1[1].Children, ns1[0])
			ns[1].Children = append(ns[1].Children, ns1[0])
		}
		return root
	} else {
		ns1 := cal1(ns[0], q, nil)
		if nil == ns[2] {
			ns1[1].Children = nodesRemoveOne(ns1[1].Children, ns1[0])
			ns1[0].Children = append(ns1[0].Children, ns[0])
			return ns1[0]
		}
		i := 0
		for _, c := range ns[2].Children {
			if c.Val == ns[0].Val {
				break
			}
			i++
		}
		ns1[1].Children = nodesRemoveOne(ns1[1].Children, ns1[0])
		ns[2].Children = nodesRemoveOne(ns[2].Children, ns[0])
		ns[2].Children = nodesInsert(ns[2].Children, i, ns1[0])
		ns1[0].Children = append(ns1[0].Children, ns[0])
		return root
	}
}
