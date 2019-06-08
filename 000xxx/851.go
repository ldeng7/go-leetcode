func loudAndRich(richer [][]int, quiet []int) []int {
	l := len(quiet)
	out, d := make([]int, l), make([]int, l)
	for i := 0; i < l; i++ {
		out[i] = -1
	}
	q := []int{}
	m := map[int][]int{}
	for _, r := range richer {
		m[r[0]] = append(m[r[0]], r[1])
		d[r[1]]++
	}
	for i, _ := range quiet {
		if d[i] == 0 {
			q = append(q, i)
		}
		out[i] = i
	}
	for 0 != len(q) {
		t := q[0]
		q = q[1:]
		for _, p := range m[t] {
			if quiet[out[p]] > quiet[out[t]] {
				out[p] = out[t]
			}
			d[p]--
			if d[p] == 0 {
				q = append(q, p)
			}
		}
	}
	return out
}
