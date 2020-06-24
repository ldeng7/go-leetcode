func minTime(n int, edges [][]int, hasApple []bool) int {
	t, v := make([]int, n), make([]bool, n)
	for _, e := range edges {
		t[e[1]] = e[0]
	}
	t[0] = -1
	o := 0
	for i, b := range hasApple {
		if !b {
			continue
		}
		for {
			j := t[i]
			if -1 == j {
				break
			}
			if !v[i] {
				v[i] = true
				o += 2
			}
			if 0 == j {
				break
			}
			i = j
		}
	}
	return o
}
