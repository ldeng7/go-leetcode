func halvesAreAlike(s string) bool {
	c, c1, l := 0, 0, len(s)
	lh := l >> 1
	for i := 0; i < lh; i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			c++
			c1++
		}
	}
	for i := lh; i < l; i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			c++
		}
	}
	return c&1 == 0 && c1<<1 == c
}
