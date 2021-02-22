import "sort"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func bestTeamScore(scores []int, ages []int) int {
	l, m := len(scores), 0
	indices := make([]int, l)
	for i, a := range ages {
		m = max(m, a)
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		a, b := indices[i], indices[j]
		sa, sb := scores[a], scores[b]
		return sa < sb || (sa == sb && ages[a] < ages[b])
	})

	t := make([]int, m+1)
	o := 0
	for _, i := range indices {
		s := 0
		for a := ages[i]; a > 0; a -= a & -a {
			s = max(s, t[a])
		}
		s += scores[i]
		o = max(o, s)
		for a := ages[i]; a <= m; a += a & -a {
			t[a] = max(t[a], s)
		}
	}
	return o
}
