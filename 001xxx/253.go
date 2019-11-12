func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	l, s := len(colsum), 0
	for _, cs := range colsum {
		s += cs
	}
	if upper+lower != s {
		return nil
	}

	o0, o1 := make([]int, l), make([]int, l)
	for i, cs := range colsum {
		switch cs {
		case 2:
			if upper >= 1 && lower >= 1 {
				o0[i], o1[i], upper, lower = 1, 1, upper-1, lower-1
			} else {
				return nil
			}
		case 1:
			if upper >= lower && upper >= 1 {
				o0[i], upper = 1, upper-1
			} else if lower >= 1 {
				o1[i], lower = 1, lower-1
			} else {
				return nil
			}
		}
	}
	return [][]int{o0, o1}
}
