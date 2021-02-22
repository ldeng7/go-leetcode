func maximumGain(s string, x int, y int) int {
	bs := []byte(s)
	cal := func(c1, c2 byte) int {
		o := 0
		bs1 := make([]byte, 1, len(bs))
		for _, c := range bs {
			if c == c2 && bs1[len(bs1)-1] == c1 {
				bs1 = bs1[:len(bs1)-1]
				o++
			} else {
				bs1 = append(bs1, c)
			}
		}
		bs = bs1
		return o
	}

	if x > y {
		return x*cal('a', 'b') + y*cal('b', 'a')
	}
	return y*cal('b', 'a') + x*cal('a', 'b')
}
