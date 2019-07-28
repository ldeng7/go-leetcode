func isIsomorphic(s string, t string) bool {
	l := len(s)
	m := make(map[byte]byte, l)
	m1 := make(map[byte]bool, l)
	for i := 0; i < l; i++ {
		cs, ct := s[i], t[i]
		if v, ok := m[cs]; !ok {
			if m1[ct] {
				return false
			}
			m[cs], m1[ct] = ct, true
		} else if v != ct {
			return false
		}
	}
	return true
}
