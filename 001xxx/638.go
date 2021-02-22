func countSubstrings(s string, t string) int {
	o, ls, lt := 0, len(s), len(t)
	for i := 0; i < ls; i++ {
		for j := 0; j < lt; j++ {
			for a, b, d := i, j, 0; a < ls && b < lt && d <= 1; a, b = a+1, b+1 {
				if s[a] != t[b] {
					d++
				}
				if d == 1 {
					o++
				}
			}
		}
	}
	return o
}
