var m [200][200]float64

func cal(a, b int) float64 {
	if a > 0 && b > 0 {
		if t := m[a][b]; t > 0 {
			return t
		}
		t := 0.25 * (cal(a-4, b) + cal(a-3, b-1) + cal(a-2, b-2) + cal(a-1, b-3))
		m[a][b] = t
		return t
	} else if a > 0 {
		return 0
	} else if b > 0 {
		return 1.0
	}
	return 0.5
}

func soupServings(N int) float64 {
	if N < 4800 {
		return cal((N+24)/25, (N+24)/25)
	}
	return 1.0
}
