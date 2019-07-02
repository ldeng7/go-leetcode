func pancakeSort(A []int) []int {
	o, l := []int{}, len(A)
	b := false
	for i := 0; i < l-1; i++ {
		if A[i] > A[i+1] {
			b = true
			break
		}
	}
	if !b {
		return []int{}
	}
	for i := 0; i < l-1; i++ {
		jm, m := 0, 0
		for j := 0; j < l-i; j++ {
			if A[j] > m {
				jm, m = j, A[j]
			}
		}
		if jm != 0 {
			for j := 0; j < (jm+1)>>1; j++ {
				A[j], A[jm-j] = A[jm-j], A[j]
			}
			o = append(o, jm+1)
		}
		e := l - i
		for j := 0; j < e>>1; j++ {
			A[j], A[e-j-1] = A[e-j-1], A[j]
		}
		o = append(o, l-i)
	}
	return o
}
