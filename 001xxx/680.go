func concatenatedBinary(n int) int {
	o := 1
	for i, j, l := 2, 4, 2; i < n+1; i++ {
		if i >= j {
			j <<= 1
			l++
		}
		o = ((o << l) + i) % 1000000007
	}
	return o
}
