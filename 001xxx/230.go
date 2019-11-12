func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func probabilityOfHeads(prob []float64, target int) float64 {
	t := make([]float64, target+1)
	t[0] = 1.0
	for i, p := range prob {
		for j := min(target, i+1); j >= 0; j-- {
			t[j] *= 1 - p
			if j != 0 {
				t[j] += t[j-1] * p
			}
		}
	}
	return t[target]
}
