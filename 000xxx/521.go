func findLUSlength(a string, b string) int {
	if a != b {
		if len(a) >= len(b) {
			return len(a)
		}
		return len(b)
	}
	return -1
}
