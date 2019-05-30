func findPermutation(s string) []int {
	out := []int{}
	for i := 0; i < len(s)+1; i++ {
		if i == len(s) || s[i] == 'I' {
			l := len(out)
			for j := i + 1; j > l; j-- {
				out = append(out, j)
			}
		}
	}
	return out
}
