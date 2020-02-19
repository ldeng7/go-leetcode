func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func popCountU8(u uint8) uint8 {
	u = (u & 0x55) + ((u >> 1) & 0x55)
	u = (u & 0x33) + ((u >> 2) & 0x33)
	u = (u & 0x0f) + ((u >> 4) & 0x0f)
	return u
}

func maxStudents(seats [][]byte) int {
	m, n := len(seats), len(seats[0])
	ar := make([]int, m)
	for i, s := range seats {
		e := 0
		for _, b := range s {
			e <<= 1
			if b == '.' {
				e |= 1
			}
		}
		ar[i] = e
	}

	n = 1 << n
	t := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		t[i] = make([]int, n)
	}
	for i := 1; i <= m; i++ {
		e := ar[i-1]
		for j := 0; j < n; j++ {
			if j&e != j || 0 != j&(j>>1) {
				continue
			}
			for k := 0; k < n; k++ {
				if 0 == j&(k>>1) && 0 == (j>>1)&k {
					t[i][j] = max(t[i][j], t[i-1][k]+int(popCountU8(uint8(j))))
				}
			}
		}
	}

	o, ar := 0, t[m]
	for _, v := range ar {
		o = max(o, v)
	}
	return o
}
