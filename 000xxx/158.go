func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	t := make([]byte, 4)
	j, k := 0, 0
	return func(buf []byte, n int) int {
		i := 0
		for i < n {
			l := n - i
			if j == k {
				k = read4(t)
				if 0 == k {
					j = k
					break
				}
				j = min(k, l)
				copy(buf[i:], t[:j])
				i += j
			} else {
				l = min(l, k-j)
				copy(buf[i:], t[j:j+l])
				i, j = i+l, j+l
			}
		}
		return i
	}
}
