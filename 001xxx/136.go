func minimumSemesters(N int, relations [][]int) int {
	degree := make([]int, N+1)
	next := make([][]int, N+1)
	for _, r := range relations {
		a, b := r[0], r[1]
		degree[b]++
		next[a] = append(next[a], b)
	}
	o, c := 0, 0
	q := [][2]int{}
	for i := 1; i <= N; i++ {
		if degree[i] == 0 {
			q = append(q, [2]int{i, 1})
		}
	}
	for 0 != len(q) {
		t := q[0]
		q = q[1:]
		o, c = t[1], c+1
		for _, ne := range next[t[0]] {
			degree[ne]--
			if degree[ne] == 0 {
				q = append(q, [2]int{ne, o + 1})
			}
		}
	}
	if c < N {
		return -1
	}
	return o
}
