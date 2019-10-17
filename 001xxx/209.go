func removeDuplicates(s string, k int) string {
	l := len(s)
	ts, cs := make([]int, 1, l+1), make([]byte, 1, l+1)
	for i := 0; i < l; i++ {
		c := s[i]
		j := len(cs) - 1
		if c == cs[j] {
			ts[j]++
			if ts[j] == k {
				ts, cs = ts[:j], cs[:j]
			}
		} else {
			ts, cs = append(ts, 1), append(cs, c)
		}
	}
	o := make([]byte, 0, l)
	for i, t := range ts {
		c := cs[i]
		for j := 0; j < t; j++ {
			o = append(o, c)
		}
	}
	return string(o)
}
