func maxNumberOfBalloons(text string) int {
	c1, c2, c3, c4, c5 := 0, 0, 0, 0, 0
	for i := 0; i < len(text); i++ {
		switch text[i] {
		case 'b':
			c1++
		case 'a':
			c2++
		case 'l':
			c3++
		case 'o':
			c4++
		case 'n':
			c5++
		}
	}
	c3, c4 = c3>>1, c4>>1
	o := c1
	if c2 < c1 {
		o = c2
	}
	if c3 < c2 {
		o = c3
	}
	if c4 < c3 {
		o = c4
	}
	if c5 < c4 {
		o = c5
	}
	return o
}
