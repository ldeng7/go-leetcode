const MOD = 1000000007

func numOfWays(n int) int {
	a, b := 6, 6
	for i := 1; i < n; i++ {
		b <<= 1
		a1, b1 := a*3+b, a*2+b
		a, b = a1%MOD, b1%MOD
	}
	return (a + b) % MOD
}
