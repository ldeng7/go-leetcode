func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxRepOpt1(text string) int {
	o, t := 1, [26][]int{}
	for i := 0; i < len(text); i++ {
		j := text[i] - 'a'
		t[j] = append(t[j], i+1)
	}
	for _, ar := range t {
		c, c1, m := 1, 0, 0
		for j := 1; j < len(ar); j++ {
			if d := ar[j] - ar[j-1]; d == 1 {
				c++
			} else if d == 2 {
				c1, c = c, 1
			} else {
				c1, c = 0, 1
			}
			m = max(m, c1+c)
		}
		if len(ar) > m {
			m++
		}
		o = max(o, m)
	}
	return o
}
