func longestRepeatingSubstring(S string) int {
	o, l := 1, len(S)
	for i := 1; o < l-i; i++ {
		c := 0
		for j := i; o-c < l-j; j++ {
			if S[j-i] == S[j] {
				c++
				if c > o {
					o = c
				}
			} else {
				c = 0
			}
		}
	}
	if o == 1 {
		return 0
	}
	return o
}
