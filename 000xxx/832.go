func flipAndInvertImage(A [][]int) [][]int {
	l := len(A)
	for _, r := range A {
		for i := 0; i < l>>1; i++ {
			r[i], r[l-i-1] = r[l-i-1], r[i]
		}
		for i := 0; i < l; i++ {
			r[i] ^= 1
		}
	}
	return A
}
