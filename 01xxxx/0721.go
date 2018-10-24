import "sort"

func find(s string, root map[string]string) string {
	if root[s] == s {
		return s
	}
	return find(root[s], root)
}

func accountsMerge(accounts [][]string) [][]string {
	out := [][]string{}
	root := map[string]string{}
	owner := map[string]string{}
	m := map[string]map[string]bool{}
	for _, a := range accounts {
		for _, s := range a[1:] {
			root[s], owner[s] = s, a[0]
		}
	}
	for _, a := range accounts {
		r := find(a[1], root)
		for _, s := range a[2:] {
			root[find(s, root)] = r
		}
	}
	for _, a := range accounts {
		for _, s := range a[1:] {
			r := find(s, root)
			if _, ok := m[r]; !ok {
				m[r] = map[string]bool{}
			}
			m[r][s] = true
		}
	}
	for r, ss := range m {
		a := []string{}
		for s, _ := range ss {
			a = append(a, s)
		}
		sort.StringSlice(a).Sort()
		a = append(a, "")
		copy(a[1:], a[:len(a)-1])
		a[0] = owner[r]
		out = append(out, a)
	}
	return out
}
