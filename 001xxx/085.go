func sumOfDigits(A []int) int {
	m := 101
	for _, n := range A {
		if n < m {
			m = n
		}
	}
	s := 0
	for ; m != 0; m /= 10 {
		s += m % 10
	}
	return (s & 1) ^ 1
}
