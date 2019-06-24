func isRobotBounded(instructions string) bool {
	x, y, d := 0, 0, 0
	for i := 0; i < 4; i++ {
		for _, c := range []byte(instructions) {
			if c == 'L' {
				d--
			} else if c == 'R' {
				d++
			} else {
				switch d % 4 {
				case 0:
					y += 1
				case 1, -3:
					x += 1
				case 2, -2:
					y -= 1
				case 3, -1:
					x -= 1
				}
			}
		}
	}
	return x == 0 && y == 0
}
