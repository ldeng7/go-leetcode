func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func strongPasswordChecker(s string) int {
	o, l := 0, len(s)
	nl, nu, nd := 1, 1, 1
	t := make([]int, l)
	for i := 0; i < l; {
		c := s[i]
		if c >= 'a' && c <= 'z' {
			nl = 0
		}
		if c >= 'A' && c <= 'Z' {
			nu = 0
		}
		if c >= '0' && c <= '9' {
			nd = 0
		}
		j := i
		for i < l && s[i] == s[j] {
			i++
		}
		t[j] = i - j
	}
	n := nl + nu + nd
	if l >= 6 {
		no, left := max(l-20, 0), 0
		o += no
		for i := 1; i < 3; i++ {
			for j := 0; j < l && no > 0; j++ {
				k := t[j]
				if k < 3 || k%3 != (i-1) {
					continue
				}
				t[j], no = k-min(no, i), no-i
			}
		}
		for i := 0; i < l; i++ {
			k := t[i]
			if k >= 3 && no > 0 {
				t[i], no = k-no, no-k+2
			}
			if t[i] >= 3 {
				left += t[i] / 3
			}
		}
		o += max(n, left)
	} else {
		d := 6 - l
		o += d + max(0, n-d)
	}
	return o
}
