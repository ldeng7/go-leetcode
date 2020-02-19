func getNoZeroIntegers(n int) []int {
	a, b, c := 0, n, 1
	for n > 0 {
		switch n % 10 {
		case 0:
			a, n = a+c, n-10
		case 1:
			if n/10 != 0 {
				a, n = a+c<<1, n-10
			} else {
				a += c
			}
		default:
			a += c
		}
		n, c = n/10, c*10
	}
	return []int{a, b - a}
}
