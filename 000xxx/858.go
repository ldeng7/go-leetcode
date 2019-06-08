func gcd(a, b int) int {
	if 0 == b {
		return a
	}
	return gcd(b, a%b)
}

func mirrorReflection(p int, q int) int {
	g := gcd(p, q)
	return 1 - (p/g)&1 + (q/g)&1
}
