func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	o, s := 0, 0
	for i := 0; i < k; i++ {
		s += calories[i]
	}
	if s < lower {
		o--
	} else if s > upper {
		o++
	}
	for i := k; i < len(calories); i++ {
		s += calories[i] - calories[i-k]
		if s < lower {
			o--
		} else if s > upper {
			o++
		}
	}
	return o
}
