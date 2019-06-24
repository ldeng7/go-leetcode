func gcd(a, b int) int {
	if 0 == b {
		return a
	}
	return gcd(b, a%b)
}

func hasGroupsSizeX(deck []int) bool {
	m := map[int]int{}
	for _, i := range deck {
		m[i]++
	}
	g := 0
	for _, c := range m {
		g = gcd(c, g)
		if g < 2 {
			return false
		}
	}
	return true
}
