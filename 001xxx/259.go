func numberOfWays(numPeople int) int {
	n := numPeople >> 1
	t := make([]int, n+1)
	t[0] = 1
	for i := 1; i <= n; i++ {
		c := 0
		for j := 1; j <= i; j++ {
			c = (c + t[j-1]*t[i-j]) % 1000000007
		}
		t[i] = c
	}
	return t[n]
}
