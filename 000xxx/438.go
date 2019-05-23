func findAnagrams(s string, p string) []int {
	bs := []byte(s)
	out := []int{}
	m := [128]int{}
	l, r, cnt, c := 0, 0, len(p), 0
	for _, b := range []byte(p) {
		m[b]++
	}
	for r < len(bs) {
		br := bs[r]
		r, c, m[br] = r+1, m[br], m[br]-1
		if c >= 1 {
			cnt--
		}
		if 0 == cnt {
			out = append(out, l)
		}
		if r-l == len(p) {
			bl := bs[l]
			l, c, m[bl] = l+1, m[bl], m[bl]+1
			if c >= 0 {
				cnt++
			}
		}
	}
	return out
}
