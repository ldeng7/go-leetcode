func bestRotation(A []int) int {
	o, l := 0, len(A)
	t := make([]int, l)
	for i := 0; i < l; i++ {
		t[(i-A[i]+1+l)%l] -= 1
	}
	for i := 1; i < l; i++ {
		t[i] += t[i-1] + 1
		if t[i] > t[o] {
			o = i
		}
	}
	return o
}
