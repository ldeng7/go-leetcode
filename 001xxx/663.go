func getSmallestString(n int, k int) string {
	o, m := make([]byte, n), 0
	for i := 0; i < n; i++ {
		n1, n2 := (n-i-1)*26, k-m
		if n2-n1 <= 1 {
			o[i] = 'a'
			m++
		} else {
			o[i] = byte(int('a') + n2 - n1 - 1)
			m += n2 - n1
		}
	}
	return string(o)
}
