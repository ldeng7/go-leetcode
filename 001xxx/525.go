func popCountU32(i int) int {
	i = (i & 0x55555555) + ((i >> 1) & 0x55555555)
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333)
	i = (i & 0x0f0f0f0f) + ((i >> 4) & 0x0f0f0f0f)
	i = (i & 0x00ff00ff) + ((i >> 8) & 0x00ff00ff)
	i = (i & 0x0000ffff) + ((i >> 16) & 0x0000ffff)
	return i
}

func numSplits(s string) int {
	l := len(s)
	ar := make([]int, l)
	for i, t := 0, 0; i < l-1; i++ {
		t |= 1 << (s[i] - 'a')
		ar[i] = popCountU32(t)
	}

	o, t := 0, 0
	for i := l - 1; i > 0; i-- {
		t |= 1 << (s[i] - 'a')
		if ar[i-1] == popCountU32(t) {
			o++
		}
	}
	return o
}
