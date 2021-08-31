func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	m := 100000
	for _, c := range baseCosts {
		m = min(m, abs(target-c))
	}
	m += target

	t := make([]bool, m)
	for _, c := range baseCosts {
		if c <= m {
			t[c-1] = true
		}
	}
	for _, c := range toppingCosts {
		for i := m - 1; i > 0; i-- {
			if j, k := i-c, i-c*2; (j >= 0 && t[j]) || (k >= 0 && t[k]) {
				t[i] = true
			}
		}
	}

	o, d := 0, 100000
	for i := 0; i < m; i++ {
		if t[i] {
			if d1 := abs(i + 1 - target); d1 < d {
				o, d = i+1, d1
			}
		}
	}
	return o
}
