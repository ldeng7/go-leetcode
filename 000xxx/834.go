func sumOfDistancesInTree(N int, edges [][]int) []int {
	t := make([][]int, N)
	ar1, ar2 := make([]int, N), make([]int, N)
	for _, e := range edges {
		a, b := e[0], e[1]
		t[a] = append(t[a], b)
		t[b] = append(t[b], a)
	}

	var cal1 func(int, int)
	cal1 = func(c, a int) {
		ar2[a] = 1
		for _, b := range t[a] {
			if b != c {
				cal1(a, b)
				ar2[a] += ar2[b]
				ar1[a] += ar1[b] + ar2[b]
			}
		}
	}

	var cal2 func(int, int)
	cal2 = func(c, a int) {
		ar1[a] = ar1[c] + N - (ar2[a] << 1)
		for _, b := range t[a] {
			if b != c {
				cal2(a, b)
			}
		}
	}

	cal1(0, 0)
	for _, b := range t[0] {
		cal2(0, b)
	}
	return ar1
}
