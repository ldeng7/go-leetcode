import "math"

func gcd(a, b int) int {
	if 0 == b {
		return a
	}
	return gcd(b, a%b)
}

func nthUglyNumber(n int, a int, b int, c int) int {
	ab, ac, bc := a*b/gcd(a, b), a*c/gcd(a, c), b*c/gcd(b, c)
	abc := a * bc / gcd(a, bc)
	l, r := 0, math.MaxInt64
	for l < r {
		m := l + (r-l)>>1
		c := m/a + m/b + m/c - m/ab - m/ac - m/bc + m/abc
		if c < n {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
