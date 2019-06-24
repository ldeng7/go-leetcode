func numDupDigitsAtMostN(N int) int {
	if N <= 10 {
		return 0
	}
	ds := []int{}
	for n := N; n > 0; n /= 10 {
		ds = append(ds, n%10)
	}
	l, c, p, j, s := len(ds), 2, 9, 9, 9
	for c < l {
		j *= p
		c, p, s = c+1, p-1, s+j
	}
	m := map[int]bool{}
	for i := l - 1; i >= 0; i-- {
		c, p, j = i, 10-l+i, ds[i]
		if i == l-1 {
			j--
		} else {
			k := 0
			for h := 0; h < j; h++ {
				if !m[h] {
					k++
				}
			}
			j = k
		}
		for c > 0 {
			j *= p
			c, p = c-1, p-1
		}
		s += j
		if m[ds[i]] {
			break
		}
		m[ds[i]] = true
	}
	if len(m) == l {
		s++
	}
	return N - s
}
