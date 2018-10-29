func isIsomorphic(s string, t string) bool {
	ms := map[byte]int{}
	mt := map[byte]int{}
	for i := 0; i < len(s); i++ {
		bs, bt := s[i], t[i]
		is, oks := ms[bs]
		it, okt := mt[bt]
		if is != it || oks != okt {
			return false
		}
		ms[bs] = i
		mt[bt] = i
	}
	return true
}
