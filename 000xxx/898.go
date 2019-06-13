func subarrayBitwiseORs(A []int) int {
	s, s1 := map[int]bool{}, map[int]bool{}
	for _, n := range A {
		s2 := map[int]bool{}
		for m, _ := range s1 {
			s2[n|m] = true
		}
		s1 = s2
		s1[n] = true
		for m, _ := range s1 {
			s[m] = true
		}
	}
	return len(s)
}
