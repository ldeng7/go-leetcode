func min(a, b float64) float64 {
	if a <= b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a <= b {
		return b
	}
	return a
}
func champagneTower(poured int, query_row int, query_glass int) float64 {
	t := [101]float64{}
	t[0] = float64(poured)
	for i := 1; i <= query_row; i++ {
		for j := i; j >= 0; j-- {
			t[j] = max(0, (t[j]-1)*0.5)
			t[j+1] += t[j]
		}
	}
	return min(1, t[query_glass])
}
