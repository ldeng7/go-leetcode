var t [1024]int

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func longestAwesome(s string) int {
	o, l := 1, len(s)
	for i := 0; i < 1024; i++ {
		t[i] = -1
	}
	for i, j := 0, 0; i < l; i++ {
		j ^= 1 << (s[i] - '0')
		t[j] = i
	}
	for i, j := 0, 0; i < l && l-i > o; i++ {
		k := t[j]
		if k > i {
			o = max(o, k-i+1)
		}
		for m := 0; m < 10; m++ {
			k = t[j^(1<<m)]
			if k > i {
				o = max(o, k-i+1)
			}
		}
		j ^= (1 << (s[i] - '0'))
	}
	return o
}
