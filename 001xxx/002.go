func commonChars(A []string) []string {
	m := [26]int{}
	for _, c := range []byte(A[0]) {
		m[c-'a']++
	}

	for _, s := range A[1:] {
		m1 := [26]int{}
		for _, c := range []byte(s) {
			if 0 != m[c-'a'] {
				m1[c-'a']++
			}
		}
		for i, n := range m {
			if n1 := m1[i]; 0 != n1 {
				if n > n1 {
					m[i] = n1
				}
			} else {
				m[i] = 0
			}
		}
	}

	out := []string{}
	for i, n := range m {
		for j := 0; j < n; j++ {
			out = append(out, string(i+'a'))
		}
	}
	return out
}
