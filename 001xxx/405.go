func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func longestDiverseString(a int, b int, c int) string {
	bs, i := make([]byte, a+b+c), 0

	rep := func(c byte, n int) {
		ie := i + n
		for ; i < ie; i++ {
			bs[i] = c
		}
	}

	var cal func(int, int, int, byte, byte, byte)
	cal = func(a, b, c int, ca, cb, cc byte) {
		if a < b {
			cal(b, a, c, cb, ca, cc)
		} else if a < c {
			cal(c, b, a, cc, cb, ca)
		} else if b < c {
			cal(a, c, b, ca, cc, cb)
		} else if 0 == b {
			rep(ca, min(a, 2))
		} else {
			m, b1 := min(a, 2), 1
			if a-m < b {
				b1 = 0
			}
			rep(ca, m)
			rep(cb, b1)
			cal(a-m, b-b1, c, ca, cb, cc)
		}
	}

	cal(a, b, c, 'a', 'b', 'c')
	return string(bs[:i])
}
