var cm = [26]int{0, 0, 1, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3}

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return nil
	}
	o, m := []string{}, map[int]int{}
	j := 0
	for i := 0; i < 9; i++ {
		j = (j << 2) | cm[s[i]-'A']
	}
	for i := 9; i < len(s); i++ {
		j = ((j << 2) & 0x000fffff) | cm[s[i]-'A']
		if m[j] == 0 {
			m[j] = 1
		} else if m[j] == 1 {
			m[j] = 2
			o = append(o, s[i-9:i+1])
		}
	}
	return o
}
