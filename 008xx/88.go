func fairCandySwap(A []int, B []int) []int {
	sa, sb := 0, 0
	m := map[int]bool{}
	for _, n := range A {
		sa += n
	}
	for _, n := range B {
		sb += n
		m[n] = true
	}
	diff := (sa - sb) >> 1
	for _, n := range A {
		if m[n-diff] {
			return []int{n, n - diff}
		}
	}
	return nil
}
