func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxUniqueSplit(s string) int {
	o, l := 0, len(s)
	m := map[string]bool{}
	var cal func(int)
	cal = func(i int) {
		if i == l {
			o = max(o, len(m))
		}
		bs := make([]byte, 0, l-i)
		for j := i; j < l; j++ {
			bs = append(bs, s[j])
			if len(m)+l-j <= o {
				return
			}
			t := string(bs)
			if m[t] {
				continue
			}
			m[t] = true
			cal(j + 1)
			delete(m, t)
		}
	}
	cal(0)
	return o
}
