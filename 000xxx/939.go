func min(a, b uint32) uint32 {
	if a <= b {
		return a
	}
	return b
}

func d(a, b uint32) uint32 {
	if a <= b {
		return b - a
	}
	return a - b
}

func minAreaRect(points [][]int) int {
	var o uint32 = 4294967295
	s := map[uint32]bool{}
	for _, p := range points {
		y1, x1 := uint32(p[0]), uint32(p[1])
		for k, _ := range s {
			y2, x2 := k>>16, k&0xffff
			if s[(y1<<16)|x2] && s[(y2<<16)|x1] {
				o = min(o, d(y2, y1)*d(x2, x1))
			}
		}
		s[(y1<<16)|x1] = true
	}
	if o == 4294967295 {
		return 0
	}
	return int(o)
}
