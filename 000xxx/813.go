func largestSumOfAverages(A []int, K int) float64 {
	l := len(A)
	s, t := make([]float64, l+1), make([]float64, l)
	for i := 0; i < l; i++ {
		s[i+1] = s[i] + float64(A[i])
	}
	for i := 0; i < l; i++ {
		t[i] = (s[l] - s[i]) / float64(l-i)
	}
	for i := 1; i < K; i++ {
		for j := 0; j < l-1; j++ {
			for k := j + 1; k < l; k++ {
				t1 := (s[k]-s[j])/float64(k-j) + t[k]
				if t1 > t[j] {
					t[j] = t1
				}
			}
		}
	}
	return t[0]
}
