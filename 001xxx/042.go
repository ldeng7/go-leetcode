func gardenNoAdj(N int, paths [][]int) []int {
	cs, out := make([][]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		cs[i] = []int{}
	}
	for _, p := range paths {
		p0, p1 := p[0]-1, p[1]-1
		if p0 > p1 {
			p0, p1 = p1, p0
		}
		cs[p1] = append(cs[p1], p0)
	}

	out[0] = 1
	for i := 1; i < N; i++ {
		gs := 0
		for _, j := range cs[i] {
			gs |= 1 << (uint(out[j]) - 1)
		}
		for g := 1; g <= 4; g++ {
			if gs&1 == 0 {
				out[i] = g
				break
			}
			gs >>= 1
		}
	}
	return out
}
