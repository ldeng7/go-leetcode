func min(a, b uint64) uint64 {
	if a <= b {
		return a
	}
	return b
}

func robot(command string, obstacles [][]int, x int, y int) bool {
	var x1, y1 uint64 = 0, 0
	set := map[uint64]bool{0: true}
	for i := 0; i < len(command); i++ {
		if command[i] == 'R' {
			x1++
		} else {
			y1++
		}
		set[(x1<<32)|y1] = true
	}

	check := func(x, y int) bool {
		m := min(uint64(x)/x1, uint64(y)/y1)
		return set[((uint64(x)-m*x1)<<32)|(uint64(y)-m*y1)]
	}

	for _, o := range obstacles {
		ox, oy := o[0], o[1]
		if ox > x || oy > y {
			continue
		}
		if check(ox, oy) {
			return false
		}
	}
	return check(x, y)
}
