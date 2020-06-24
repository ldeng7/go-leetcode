type Node struct {
	Val  int
	Next *Node
}

func processQueries(queries []int, m int) []int {
	h := &Node{1, nil}
	n := h
	for i := 2; i <= m; i++ {
		n.Next = &Node{i, nil}
		n = n.Next
	}
	o := make([]int, 0, len(queries))
	for _, q := range queries {
		if h.Val == q {
			o = append(o, 0)
			continue
		}
		c, n := 1, h
		for n.Next.Val != q {
			c, n = c+1, n.Next
		}
		o = append(o, c)
		t := n.Next
		n.Next, t.Next, h = n.Next.Next, h, t
	}
	return o
}
