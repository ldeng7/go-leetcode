func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func new21Game(N int, K int, W int) float64 {
	if K == 0 || N >= K+W {
		return 1.0
	}
	s := make([]float64, K+W)
	s[0] = 1.0
	for i := 1; i < K+W; i++ {
		t := min(i-1, K-1)
		if i <= W {
			s[i] = s[i-1] + s[t]/float64(W)
		} else {
			s[i] = s[i-1] + (s[t]-s[i-W-1])/float64(W)
		}
	}
	return (s[N] - s[K-1]) / (s[K+W-1] - s[K-1])
}
