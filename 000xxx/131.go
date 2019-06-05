func partition(s string) [][]string {
	l := len(s)
	out := [][]string{}
	ar := []string{}
	t := make([][]bool, l)
	for i := 0; i < l; i++ {
		t[i] = make([]bool, l)
	}
	for i := 0; i < l; i++ {
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && (i-j <= 2 || t[j+1][i-1]) {
				t[j][i] = true
			}
		}
	}

	var cal func(int)
	cal = func(b int) {
		if b == l {
			ar1 := make([]string, len(ar))
			copy(ar1, ar)
			out = append(out, ar1)
			return
		}
		for i := b; i < l; i++ {
			if !t[b][i] {
				continue
			}
			ar = append(ar, s[b:i+1])
			cal(i + 1)
			ar = ar[:len(ar)-1]
		}
	}

	cal(0)
	return out
}
