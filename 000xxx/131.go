func partition(s string) [][]string {
	o, l, ar := [][]string{}, len(s), []string{}
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
			o = append(o, ar1)
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
	return o
}
