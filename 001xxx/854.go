func maximumPopulation(logs [][]int) int {
	o, m := -1, 0
	for i := 1950; i <= 2050; i++ {
		p := 0
		for _, l := range logs {
			if i >= l[0] && i < l[1] {
				p++
			}
		}
		if p > m {
			o, m = i, p
		}
	}
	return o
}
