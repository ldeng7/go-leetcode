func cal(n, d int) int {
	c := 0
	for i := 1; n/i != 0; i *= 10 {
		k := n / i
		h := k / 10
		if d == 0 {
			if h != 0 {
				h--
			} else {
				break
			}
		}
		c += h * i
		j := k % 10
		if j > d {
			c += i
		} else if j == d {
			c += n - k*i + 1
		}
	}
	return c
}

func digitsCount(d int, low int, high int) int {
	return cal(high, d) - cal(low-1, d)
}
