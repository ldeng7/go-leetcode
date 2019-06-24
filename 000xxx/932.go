func beautifulArray(N int) []int {
	m, k := N-1, 1
	for m > 1 {
		m, k = m>>1, k<<1
	}
	o := make([]int, N)
	o[0] = 1
	i, pi := 1, 1
	for i < N {
		for j := 0; j < pi; j++ {
			if o[j]+k <= N {
				o[i], i = o[j]+k, i+1
			}
		}
		pi, k = i, k>>1
	}
	return o
}
