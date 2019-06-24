func sumSubarrayMins(A []int) int {
	o, n := 0, len(A)
	for i, a := range A {
		l, r := i-1, i+1
		for ; l >= 0 && A[l] > a; l-- {
		}
		for ; r < n && A[r] >= a; r++ {
		}
		o += (i - l) * (r - i) * a
	}
	return o % 1000000007
}
