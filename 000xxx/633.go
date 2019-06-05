func judgeSquareSum(c int) bool {
	for i := 2; i*i <= c; i++ {
		if c%i != 0 {
			continue
		}
		cnt := 0
		for c%i == 0 {
			cnt++
			c /= i
		}
		if i%4 == 3 && cnt&1 != 0 {
			return false
		}
	}
	return c%4 != 3
}
