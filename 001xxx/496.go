func isPathCrossing(path string) bool {
	l := len(path)
	var x, y int16
	m := map[uint32]bool{0: true}
	for i := 0; i < l; i++ {
		switch path[i] {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		case 'W':
			x--
		}
		k := (uint32(uint16(y)) << 16) | uint32(uint16(x))
		if m[k] {
			return true
		}
		m[k] = true
	}
	return false
}
