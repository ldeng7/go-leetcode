func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func guardCastle(grid []string) int {
	l := len(grid[0]) + 2
	g0, g1, g0t, g1t := make([]byte, l), make([]byte, l), make([]byte, l), make([]byte, l)
	copy(g0[1:], grid[0])
	copy(g1[1:], grid[1])
	g0[0], g0[l-1], g1[0], g1[l-1] = '.', '.', '.', '.'
	copy(g0t, g0)
	copy(g1t, g1)

	cal1 := func(i, p int) int {
		for j := p + 1; j <= i; j++ {
			if (g0[j] == '#' && (g1[j] == '#' || g1[j-1] == '#')) ||
				(g1[j] == '#' && (g0[j] == '#' || g0[j-1] == '#')) {
				return 0
			}
		}
		for j := i; j >= p; j-- {
			if g0[j] == '#' {
				if j < i && g1[j+1] == '.' {
					g1[j+1] = '#'
				}
				return 1
			}
			if g1[j] == '#' {
				if j < i && g0[j+1] == '.' {
					g0[j+1] = '#'
				}
				return 1
			}
		}
		if g0[i] == '.' {
			g0[i] = '#'
		}
		if g1[i] == '.' {
			g1[i] = '#'
		}
		return 2
	}

	cal := func() int {
		for i := 1; i < l; i++ {
			if g0[i]+g1[i] == 3 || g0[i-1]+g0[i] == 3 || g1[i-1]+g1[i] == 3 {
				return 1e5
			}
		}
		o, p, t := 0, 0, -1
		for i := 1; i < l; i++ {
			if g0[i] == 1 || g1[i] == 1 {
				if t == 2 {
					o += cal1(i, p)
				}
				p, t = i, 1
			} else if g0[i] == 2 || g1[i] == 2 {
				if t == 1 {
					o += cal1(i, p)
				}
				p, t = i, 2
			}
		}
		return o
	}

	for i := 1; i < l-1; i++ {
		switch g0[i] {
		case 'S', 'P':
			g0[i] = 1
		case 'C':
			g0[i] = 2
		}
		switch g1[i] {
		case 'S', 'P':
			g1[i] = 1
		case 'C':
			g1[i] = 2
		}
	}
	o := cal()

	g0, g1 = g0t, g1t
	for i := 1; i < l-1; i++ {
		switch g0[i] {
		case 'S':
			g0[i] = 1
		case 'C', 'P':
			g0[i] = 2
		}
		switch g1[i] {
		case 'S':
			g1[i] = 1
		case 'C', 'P':
			g1[i] = 2
		}
	}
	o = min(o, cal())

	if o == 1e5 {
		return -1
	}
	return o
}
