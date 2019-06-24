func gcd(a, b int) int {
	if 0 == b {
		return a
	}
	return gcd(b, a%b)
}

func canMeasureWater(x int, y int, z int) bool {
	return z == 0 || (x+y >= z && z%gcd(x, y) == 0)
}
