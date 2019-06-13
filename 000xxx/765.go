func minSwapsCouples(row []int) int {
	m := map[int]int{}
	var cal func(a, b int)
	cal = func(a, b int) {
		if a > b {
			a, b = b, a
		}
		if a == b {
			return
		}
		if _, ok := m[a]; ok {
			cal(m[a], b)
		} else {
			m[a] = b
		}
	}

	for i := 0; i < len(row); i += 2 {
		cal(row[i]>>1, row[i+1]>>1)
	}
	return len(m)
}
