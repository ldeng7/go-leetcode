var gens = [4]byte{'A', 'C', 'G', 'T'}

func minMutation(start string, end string, bank []string) int {
	if 0 == len(bank) {
		return -1
	}
	s, v := map[string]bool{}, map[string]bool{}
	for _, t := range bank {
		s[t] = true
	}
	q, l := []string{start}, 0
	for 0 != len(q) {
		for i := len(q); i > 0; i-- {
			st := q[0]
			t := []byte(st)
			q = q[1:]
			if st == end {
				return l
			}
			for j := 0; j < len(t); j++ {
				c0 := t[j]
				for _, c := range gens {
					t[j] = c
					st := string(t)
					if s[st] && !v[st] {
						v[st] = true
						q = append(q, st)
					}
				}
				t[j] = c0
			}
		}
		l++
	}
	return -1
}
