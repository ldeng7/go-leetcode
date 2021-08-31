func longestNiceSubstring(s string) string {
	n, bs := 0, [26]bool{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < 'a' {
			c += 32
		}
		c -= 'a'
		if !bs[c] {
			bs[c] = true
			n++
		}
	}

	b, e := 0, -1
	for m := 1; m <= n; m++ {
		ls, us := [26]int{}, [26]int{}
		b1, e1 := 0, -1
		for l, r, n := 0, 0, 0; r < len(s); r++ {
			c := s[r]
			if c >= 'a' {
				c -= 'a'
				ls[c]++
			} else {
				c -= 'A'
				us[c]++
			}
			if ls[c]+us[c] == 1 {
				n++
			}

			for ; n > m; l++ {
				c = s[l]
				if c >= 'a' {
					c -= 'a'
					if ls[c]+us[c] == 1 {
						n--
					}
					ls[c]--
				} else {
					c -= 'A'
					if ls[c]+us[c] == 1 {
						n--
					}
					us[c]--
				}
			}

			s := 0
			for i := 0; i < 26; i++ {
				if ls[i] > 0 && us[i] > 0 {
					s++
				}
			}
			if s == n && r-l > e1-b1 {
				b1, e1 = l, r
			}
		}
		if e1-b1 > e-b {
			b, e = b1, e1
		}
	}
	return s[b : e+1]
}
