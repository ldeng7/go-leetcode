func maxRotateFunction(A []int) int {
	o, j, s, l := 0, 0, 0, len(A)
	for i := 0; i < l; i++ {
		s += A[i]
		j += i * A[i]
	}
	o = j
	for i := 1; i < l; i++ {
		j += s - l*A[l-i]
		if j > o {
			o = j
		}
	}
	return o
}
