func addToArrayForm(A []int, K int) []int {
	o := []int{}
	for i := len(A) - 1; K != 0 || i >= 0; i-- {
		if i >= 0 {
			K += A[i]
		}
		o, K = append(o, K%10), K/10
	}
	l := len(o)
	for i := 0; i < l>>1; i++ {
		o[i], o[l-i-1] = o[l-i-1], o[i]
	}
	return o
}
