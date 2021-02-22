func differByOne(dict []string) bool {
	if len(dict) < 2 {
		return false
	}
	set := map[string]bool{}
	for _, s := range dict {
		bs := []byte(s)
		for i := 0; i < len(s); i++ {
			bs[i] = 0
			s1 := string(bs)
			if set[s1] {
				return true
			}
			set[s1] = true
			bs[i] = s[i]
		}
	}
	return false
}
