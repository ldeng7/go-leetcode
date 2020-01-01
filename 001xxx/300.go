func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func findBestValue(arr []int, target int) int {
	l := len(arr)
	av, s, c := target/l, 0, 0
	for _, n := range arr {
		if av >= n {
			s, c = s+n, c+1
		}
	}
	o := (target - s) / (l - c)
	if abs(s+o*(l-c)-target) > abs(s+(o+1)*(l-c)-target) {
		o++
	}
	return o
}
