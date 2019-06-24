func reverseStr(s string, k int) string {
	l, bs := len(s), []byte(s)
	for i := 0; i < l; i += k << 1 {
		j1 := i + k
		if j1 > l {
			j1 = l
		}
		jm := i + (j1-i)>>1
		for j := i; j < jm; j++ {
			j1--
			bs[j], bs[j1] = bs[j1], bs[j]
		}
	}
	return string(bs)
}
