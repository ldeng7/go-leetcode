func restoreString(s string, indices []int) string {
	bs := make([]byte, len(s))
	for i, j := range indices {
		bs[j] = s[i]
	}
	return string(bs)
}
