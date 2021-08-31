func isThree(n int) bool {
	c := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if i != n/i {
				c += 2
			} else {
				c += 1
			}
		}
	}
	return c == 3
}
