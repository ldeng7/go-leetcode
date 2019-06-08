func minmaxGasDist(stations []int, K int) float64 {
	var l, r float64 = 0, 1e8
	for r-l > 1e-6 {
		c, m := 0, l+(r-l)>>1
		for i := 0; i < len(stations)-1; i++ {
			c += int(float64(stations[i+1]-stations[i]) / m)
		}
		if c <= K {
			r = m
		} else {
			l = m
		}
	}
	return l
}
