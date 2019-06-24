func catMouseGame(graph [][]int) int {
	l := byte(len(graph))
	t := make([][][2]byte, l)
	for i := byte(0); i < l; i++ {
		li := byte(len(graph[i]))
		ar := make([][2]byte, l)
		for j := byte(0); j < l; j++ {
			ar[j][0], ar[j][1] = li, byte(len(graph[j]))
		}
		t[i] = ar
	}
	q := [][4]byte{}
	for i := byte(1); i < l; i++ {
		q = append(q, [4]byte{i, i, 0, 1}, [4]byte{i, i, 1, 0}, [4]byte{0, i, 1, 1}, [4]byte{i, 0, 0, 0})
		t[i][i][0], t[i][i][1], t[0][i][0], t[0][i][1], t[i][0][0], t[i][0][1] = 0, 0, 0, 0, 0, 0
	}
	for 0 != len(q) {
		l := len(q)
		for i := 0; i < l; i++ {
			h := q[0]
			p := h[2] ^ 1
			h[2], h[3], q = p, h[3]^1, q[1:]
			g := graph[h[p]]
			for _, j := range g {
				h[p] = byte(j)
				b := &(t[h[0]][h[1]][p])
				if *b != 0 {
					if h[3] == 1 {
						(*b)--
					}
					if h[3] == 0 || *b == 0 {
						*b, q = 0, append(q, h)
						if p == 0 && h[0] == 1 && h[1] == 2 {
							return int(h[3] + 1)
						}
					}
				}
			}
		}
	}
	return 0
}
