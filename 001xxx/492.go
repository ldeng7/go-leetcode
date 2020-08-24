func kthFactor(n int, k int) int {
	var i, c int
	for i = 1; i*i <= n; i++ {
		if n%i == 0 {
			c++
			if c == k {
				return i
			}
		}
	}
	i--
	if i*i == n {
		i--
	}
	for ; i > 0; i-- {
		if n%i == 0 {
			c++
			if c == k {
				return n / i
			}
		}
	}
	return -1
}
