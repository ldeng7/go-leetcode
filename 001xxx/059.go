type Node struct {
	val  int
	next *Node
}

func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
	nodes, bs := make([]*Node, n), make([]bool, n)
	for i := 0; i < n; i++ {
		nodes[i] = &Node{val: i}
	}
	for _, e := range edges {
		n1 := nodes[e[0]]
		n := &Node{val: e[1], next: n1.next}
		n1.next = n
	}

	var cal func(int) bool
	cal = func(i int) bool {
		if nil == nodes[i].next {
			return i == destination
		}
		bs[i] = true
		ne := nodes[i].next
		for nil != ne {
			if bs[ne.val] || !cal(ne.val) {
				return false
			}
			ne = ne.next
		}
		bs[i] = false
		return true
	}
	return cal(source)
}

