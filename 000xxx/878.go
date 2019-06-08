func gcd(a, b int) int {
	if 0 == b {
		return a
	}
	return gcd(b, a%b)
}

func nthMagicalNumber(N int, A int, B int) int {
	l, r, g := 1, int(1e18), gcd(A, B)
	lcm := A * B / g
	out := 0
	for l <= r {
		m := l + (r-l)>>1
		if m/A+m/B-m/lcm >= N {
			out, r = m, m-1
		} else {
			l = m + 1
		}
	}
	return out % (1e9 + 7)
}
