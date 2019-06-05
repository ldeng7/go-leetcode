func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	out := make([][]int, R*C)
	out[0] = []int{r0, c0}
	i, l := 1, R*C
	cal := func(a, b int) {
		y, x := r0+a, c0+b
		if y >= 0 && y < R && x >= 0 && x < C {
			out[i] = []int{y, x}
			i++
		}
	}
	cal(1, 0)
	cal(-1, 0)
	cal(0, 1)
	cal(0, -1)
	for d := 2; i < l; d++ {
		cal(d, 0)
		cal(-d, 0)
		cal(0, d)
		cal(0, -d)
		for a := 1; a <= (d-1)>>1; a++ {
			b := d - a
			cal(a, b)
			cal(b, a)
			cal(-a, b)
			cal(-b, a)
			cal(a, -b)
			cal(b, -a)
			cal(-a, -b)
			cal(-b, -a)
		}
		if d&1 == 0 {
			a := d >> 1
			cal(a, a)
			cal(-a, a)
			cal(a, -a)
			cal(-a, -a)
		}
	}
	return out
}
