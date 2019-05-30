func smallestFactorization(a int) int {
	if a < 10 {
		return a
	}
	out, cnt := 0, 1
	for i := 9; i >= 2; i-- {
		for a%i == 0 {
			out += cnt * i
			if out > 2147483647 {
				return 0
			}
			a /= i
			cnt *= 10
		}
	}
	if a == 1 {
		return out
	}
	return 0
}
