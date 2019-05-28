func cal(m, n int, u, i0, j0 uint) int {
	if 0 == n {
		return 1
	}
	o := 0
	if m <= 0 {
		o++
	}
	var i, j uint
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			i1, j1, u1 := i0+i, j0+j, u|(1<<(i*3+j))
			if u1 > u && (i1&1 == 1 || j1&1 == 1 || u1&(1<<(i1/2*3+j1/2)) != 0) {
				o += cal(m-1, n-1, u1, i, j)
			}
		}
	}
	return o
}

func numberOfPatterns(m int, n int) int {
	return cal(m, n, 0, 1, 1)
}
