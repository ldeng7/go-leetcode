func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func cost(a int) int {
	if a > 0 {
		return a
	}
	return 2
}

func leastOpsExpressTarget(x int, target int) int {
	t := map[complex128]int{}
	var cal func(int, int) int
	cal = func(i, tg int) int {
		k := complex(float64(i), float64(tg))
		if o, ok := t[k]; ok {
			return o
		}
		var o int
		if tg == 0 {
			o = 0
		} else if tg == 1 {
			o = cost(i)
		} else if i >= 39 {
			o = tg + 1
		} else {
			m, r, c := tg/x, tg%x, cost(i)
			o = min(r*c+cal(i+1, m), (x-r)*c+cal(i+1, m+1))
		}
		t[k] = o
		return o
	}
	return cal(0, target) - 1
}
