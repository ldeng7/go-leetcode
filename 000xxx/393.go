func validUtf8(data []int) bool {
	c := 0
	for _, b := range data {
		if c == 0 {
			if b&0xe0 == 0xc0 {
				c = 1
			} else if b&0xf0 == 0xe0 {
				c = 2
			} else if b&0xf8 == 0xf0 {
				c = 3
			} else if b&0x80 == 0x80 {
				return false
			}
		} else {
			if b&0xc0 != 0x80 {
				return false
			}
			c--
		}
	}
	return c == 0
}
