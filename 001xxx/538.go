func guessMajority(reader *ArrayReader) int {
	o, l := 0, reader.length()
	c1, c2, v := 1, 0, reader.query(0, 1, 2, 3)
	for i := 4; i < l; i++ {
		if reader.query(0, 1, 2, i) == v {
			c1++
		} else {
			o = i
			c2++
		}
	}

	v = reader.query(0, 1, 2, 4)
	if reader.query(1, 2, 3, 4) == v {
		c1++
	} else {
		c2++
		o = 0
	}
	if reader.query(0, 1, 3, 4) == v {
		c1++
	} else {
		c2++
		o = 2
	}
	if reader.query(0, 2, 3, 4) == v {
		c1++
	} else {
		c2++
		o = 1
	}

	if c1 == c2 {
		return -1
	} else if c1 > c2 {
		return 3
	}
	return o
}
