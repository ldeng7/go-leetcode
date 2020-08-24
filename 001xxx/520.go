func maxNumOfSubstrings(s string) []string {
	l := len(s)
	ar1, ar2 := [26]int{}, [26]int{}
	for i := 0; i < 26; i++ {
		ar1[i], ar2[i] = -1, -1
	}
	for i, j := 0, l-1; i < l; i, j = i+1, j-1 {
		if c := s[i] - 'a'; ar1[c] == -1 {
			ar1[c] = i
		}
		if c := s[j] - 'a'; ar2[c] == -1 {
			ar2[c] = j
		}
	}
	o, t1, t2 := make([]string, 0, l), [26]int{}, [26]int{}
	for i := 0; i < l; i++ {
		c := s[i] - 'a'
		t2[c] = i
		if ar1[c] == i {
			t1[c] = i
		}
		if ar2[c] != i || ar1[c] != t1[c] {
			continue
		}
		b1, b2, k := true, true, ar1[c]
		for b2 {
			b2 = false
			for j := byte(0); j < 26; j++ {
				if j == c || t2[j] <= k {
					continue
				}
				if t1[j] != ar1[j] || t2[j] != ar2[j] {
					b1 = false
					break
				} else if t1[j] < k {
					k, b2 = t1[j], true
				}
			}
		}
		if b1 {
			for i := 0; i < 26; i++ {
				t1[i] = -1
			}
			o = append(o, s[k:i+1])
		}
	}
	return o
}
