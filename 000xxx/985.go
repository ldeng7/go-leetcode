func sumEvenAfterQueries(A []int, queries [][]int) []int {
	o, s := make([]int, len(queries)), 0
	for _, a := range A {
		if a&1 == 0 {
			s += a
		}
	}
	for i, q := range queries {
		a := A[q[1]]
		if a&1 == 0 {
			s -= a
		}
		a += q[0]
		A[q[1]] = a
		if a&1 == 0 {
			s += a
		}
		o[i] = s
	}
	return o
}
