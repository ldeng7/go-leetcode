func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
	t := make([][]bool, n)
	for i := 0; i < n; i++ {
		t[i] = make([]bool, n)
	}
	for _, p := range prerequisites {
		t[p[0]][p[1]] = true
	}

	for i := 0; i < n; i++ {
		ari := t[i]
		for j := 0; j < n; j++ {
			arj := t[j]
			for k := 0; k < n; k++ {
				arj[k] = arj[k] || (arj[i] && ari[k])
			}
		}
	}

	o := make([]bool, len(queries))
	for i, q := range queries {
		o[i] = t[q[0]][q[1]]
	}
	return o
}
