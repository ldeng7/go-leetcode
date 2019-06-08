func isValidSerialization(preorder string) bool {
	c, bs := 1, []byte(preorder)
	bs = append(bs, ',')
	preorder += ","
	for i, b := range bs {
		if b != ',' {
			continue
		}
		c--
		if c < 0 {
			return false
		}
		if bs[i-1] != '#' {
			c += 2
		}
	}
	return c == 0
}
