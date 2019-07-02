func minMalwareSpread(graph [][]int, initial []int) int {
	o, l, mx := 0, len(graph), 0
	t, m, s := make([]int, l), map[int]int{}, map[int]bool{}
	for i := 0; i < l; i++ {
		t[i] = -1
	}
	for _, i := range initial {
		s[i] = true
	}

	var cal func(int, int)
	cal = func(i, j int) {
		if t[i] == -1 {
			t[i] = j
		} else {
			t[i] = -2
		}
		for k := 0; k < l; k++ {
			if graph[i][k] != 0 && !s[k] && t[k] != j && t[k] != -2 {
				cal(k, j)
			}
		}
	}

	for _, i := range initial {
		cal(i, i)
	}
	for _, i := range t {
		if i >= 0 {
			m[i]++
			if m[i] > mx {
				mx, o = m[i], i
			} else if m[i] == mx && i < o {
				o = i
			}
		}
	}
	if len(m) == 0 {
		o = initial[0]
		for i := 1; i < len(initial); i++ {
			if t := initial[i]; t < o {
				o = t
			}
		}
	}
	return o
}
