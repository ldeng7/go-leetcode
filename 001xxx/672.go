func maximumWealth(accounts [][]int) int {
	o := 0
	for _, ar := range accounts {
		s := 0
		for _, a := range ar {
			s += a
		}
		if s > o {
			o = s
		}
	}
	return o
}
