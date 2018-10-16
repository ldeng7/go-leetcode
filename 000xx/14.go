func longestCommonPrefix(strs []string) string {
	if 0 == len(strs) {
		return ""
	}
	cs := []byte{}
	l := len(strs[0])
	for _, s := range strs {
		if len(s) < l {
			l = len(s)
		}
	}
	for i := 0; i < l; i++ {
		c := strs[0][i]
		j := 1
		for ; j < len(strs); j++ {
			if strs[j][i] != c {
				break
			}
		}
		if j == len(strs) {
			cs = append(cs, c)
		} else {
			break
		}
	}
	return string(cs)
}
