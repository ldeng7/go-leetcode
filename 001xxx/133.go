func largestUniqueNumber(A []int) int {
	m := map[int]int{}
	for _, a := range A {
		m[a]++
	}
	o := -1
	for a, n := range m {
		if n == 1 && a > o {
			o = a
		}
	}
	return o
}
