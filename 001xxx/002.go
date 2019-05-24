func commonChars(A []string) []string {
	m := &[26]int{}
	for _, c := range []byte(A[0]) {
		(*m)[c-'a']++
	}

	for _, s := range A[1:] {
		m1 := &[26]int{}
		for _, c := range []byte(s) {
			c -= 'a'
			if 0 != (*m)[c] {
				(*m)[c]--
				(*m1)[c]++
			}
		}
		m = m1
	}

	out := []string{}
	for i, n := range *m {
		for j := 0; j < n; j++ {
			out = append(out, string(i+'a'))
		}
	}
	return out
}
