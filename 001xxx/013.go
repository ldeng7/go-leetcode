func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, a := range A {
		sum += a
	}
	if sum%3 != 0 {
		return false
	}
	sum /= 3

	s, c := 0, 0
	for _, a := range A {
		s += a
		if s == sum {
			s, c = 0, c+1
		}
	}
	return c == 3 && s == 0
}
