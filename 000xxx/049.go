func groupAnagrams(strs []string) [][]string {
	o, m := [][]string{}, map[string]int{}
	for _, s := range strs {
		ar := [26]byte{}
		for i := 0; i < len(s); i++ {
			ar[s[i]-'a']++
		}
		k := string(ar[:])
		if i, ok := m[k]; ok {
			o[i] = append(o[i], s)
		} else {
			o = append(o, []string{s})
			m[k] = len(o) - 1
		}
	}
	return o
}
