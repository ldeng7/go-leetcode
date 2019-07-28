func numEquivDominoPairs(dominoes [][]int) int {
	m := map[uint8]int{}
	for _, d := range dominoes {
		a, b := uint8(d[0]), uint8(d[1])
		if a > b {
			a, b = b, a
		}
		m[(a<<4)|b]++
	}
	o := 0
	for _, c := range m {
		if c > 2 {
			o += c * (c - 1) >> 1
		} else if c == 2 {
			o++
		}
	}
	return o
}
