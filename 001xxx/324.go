func printVertically(s string) []string {
	l := len(s)
	indices := make([]int, 1, (l>>1)+2)
	lm, bp := 0, 0
	for i := 0; i < l; i++ {
		if s[i] == ' ' {
			indices = append(indices, i+1)
			if l1 := i - bp; l1 > lm {
				lm = l1
			}
			bp, i = i+1, i+1
		}
	}
	indices = append(indices, l+1)
	if l1 := l - bp; l1 > lm {
		lm = l1
	}

	n, lm := lm, len(indices)-1
	ar := make([][]byte, n)
	ls := make([]int, n)
	for i, j := lm-1, 0; i >= 0 && j < n; i-- {
		l := indices[i+1] - 1 - indices[i]
		for ; j < l; j++ {
			ls[j] = i + 1
		}
	}
	for i := 0; i < n; i++ {
		l := ls[i]
		bs := make([]byte, l)
		for j := 0; j < l; j++ {
			bs[j] = ' '
		}
		ar[i] = bs
	}

	for i := 0; i < lm; i++ {
		e := indices[i+1] - 1
		for j, k := indices[i], 0; j < e; j, k = j+1, k+1 {
			ar[k][i] = s[j]
		}
	}
	o := make([]string, n)
	for i := 0; i < n; i++ {
		o[i] = string(ar[i])
	}
	return o
}
