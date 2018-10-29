func findRepeatedDnaSequences(s string) []string {
	m := map[string]bool{}
	mo := map[string]bool{}
	for i := 0; i < len(s)-9; i++ {
		t := s[i : i+10]
		if m[t] {
			mo[t] = true
		} else {
			m[t] = true
		}
	}
	out := []string{}
	for t, _ := range mo {
		out = append(out, t)
	}
	return out
}
