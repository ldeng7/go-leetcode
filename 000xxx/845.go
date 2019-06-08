func longestMountain(A []int) int {
	o, u, d, l := 0, 0, 0, len(A)
	for i := 1; i < l; i++ {
		if (d != 0 && A[i-1] < A[i]) || (A[i-1] == A[i]) {
			u, d = 0, 0
		}
		if A[i-1] < A[i] {
			u++
		}
		if A[i-1] > A[i] {
			d++
		}
		if u > 0 && d > 0 && u+d+1 > o {
			o = u + d + 1
		}
	}
	return o
}
