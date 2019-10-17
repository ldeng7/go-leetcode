func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	l := n + m
	g, items := make([][]int, l+1), make([][]int, m)
	for i, gr := range group {
		gid := i
		if gr != -1 {
			gid = n + gr
		}
		for _, j := range beforeItems[i] {
			grj := group[j]
			if grj == -1 {
				g[gid] = append(g[gid], j)
			} else if grj == gr {
				g[i] = append(g[i], j)
			} else {
				g[gid] = append(g[gid], n+grj)
			}
		}
	}

	vis, p := make([]bool, l), make([]int, l)
	for i := 0; i < l; i++ {
		p[i] = -1
	}
	var cal func(int, *[]int) bool
	cal = func(i int, v *[]int) bool {
		vis[i] = true
		for _, id := range g[i] {
			if !vis[id] {
				p[id] = i
				if !cal(id, v) {
					return false
				}
			} else {
				for j := p[i]; j != -1; j = p[j] {
					if j == id {
						return false
					}
				}
			}
		}
		*v = append(*v, i)
		return true
	}

	o := []int{}
	for i, gr := range group {
		if gr != -1 && !vis[i] && !cal(i, &items[gr]) {
			return o
		}
	}
	order := []int{}
	for i := 0; i < l; i++ {
		if ((i < n && group[i] == -1) || i >= n) && !vis[i] && !cal(i, &order) {
			return o
		}
	}
	for _, ord := range order {
		if ord >= n {
			o = append(o, items[ord-n]...)
		} else {
			o = append(o, ord)
		}
	}
	return o
}
