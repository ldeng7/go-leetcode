func longestPrefix(s string) string {
	l := len(s)
	ar := make([]int, l)
	for i, j := 1, 0; i < l; {
		if s[i] == s[j] {
			ar[i] = j + 1
			i, j = i+1, j+1
		} else if j != 0 {
			j = ar[j-1]
		} else {
			i++
		}
	}
	return s[:ar[l-1]]
}
