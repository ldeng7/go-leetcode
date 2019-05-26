var solution = func(read4 func([]byte) int) func([]byte, int) int {
	return func(buf []byte, n int) int {
		i := 0
		t := make([]byte, 4)
		for i < n {
			l := n - i
			k := read4(t)
			if 0 == k {
				break
			}
			if k < l {
				l = k
			}
			copy(buf[i:], t[:l])
			i += l
		}
		return i
	}
}
