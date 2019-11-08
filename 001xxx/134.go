func isArmstrong(N int) bool {
	n := N
	ds, k := make([]int, 0, 9), 0
	for n != 0 {
		ds = append(ds, n%10)
		n /= 10
		k++
	}
	s := 0
	for _, d := range ds {
		m := 1
		for i := 0; i < k; i++ {
			m *= d
		}
		s += m
	}
	return s == N
}
