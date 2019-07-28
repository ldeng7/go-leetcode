func isHappy(n int) bool {
	m := map[int]bool{}
	for n != 1 {
		t := 0
		for n > 0 {
			t += (n % 10) * (n % 10)
			n /= 10
		}
		if m[t] {
			break
		}
		m[t], n = true, t
	}
	return n == 1
}
