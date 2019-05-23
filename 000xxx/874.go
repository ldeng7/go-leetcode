func robotSim(commands []int, obstacles [][]int) int {
	mxy := map[int]map[int]bool{}
	myx := map[int]map[int]bool{}
	for _, o := range obstacles {
		ox, oy := o[0], o[1]
		if mx, ok := mxy[ox]; ok {
			mx[oy] = true
		} else {
			mxy[ox] = map[int]bool{oy: true}
		}
		if my, ok := myx[oy]; ok {
			my[ox] = true
		} else {
			myx[oy] = map[int]bool{ox: true}
		}
	}

	y, x, dy, dx, out := 0, 0, 1, 0, 0
	for _, c := range commands {
		if c == -1 {
			dy, dx = -dx, dy
		} else if c == -2 {
			dy, dx = dx, -dy
		} else if dy != 0 {
			if mx, ok := mxy[x]; ok {
				i := 0
				for ; i < c; i++ {
					if _, ok := mx[y+(i+1)*dy]; ok {
						break
					}
				}
				y += i * dy
			} else {
				y += c * dy
			}
		} else {
			if my, ok := myx[y]; ok {
				i := 0
				for ; i < c; i++ {
					if _, ok := my[x+(i+1)*dx]; ok {
						break
					}
				}
				x += i * dx
			} else {
				x += c * dx
			}
		}
		if d := x*x + y*y; d > out {
			out = d
		}
	}
	return out
}
