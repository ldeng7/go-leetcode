func findLongestWord(s string, d []string) string {
	var o string
	for _, t := range d {
		i := 0
		for j := 0; j < len(s); j++ {
			if i < len(t) && s[j] == t[i] {
				i++
			}
		}
		if i == len(t) && len(t) >= len(o) {
			if len(t) > len(o) || t < o {
				o = t
			}
		}
	}
	return o
}
