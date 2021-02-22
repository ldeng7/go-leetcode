const MOD = 1000000007

func p(a, b int) int {
	if b >= 2 {
		c := p(a, b/2)
		c = (c * c) % MOD
		if b&1 == 1 {
			return (a * c) % MOD
		}
		return c
	} else if b == 1 {
		return a
	}
	return 1
}

func numberOfSets(n int, k int) int {
	ar := make([]int, 1, n+k+1)
	ar[0] = 1
	a := 1
	for i := 1; i <= n+k; i++ {
		a = (a * i) % MOD
		ar = append(ar, a)
	}
	a, b := n+k-1, k*2
	return (ar[a] * p(ar[b], MOD-2) % MOD) * p(ar[a-b], MOD-2) % MOD
}
