func countSteppingNumbers(low int, high int) []int {
	if high == 0 {
		return []int{0}
	}
	o := []int{}
	ll, lh, t := 0, 0, low
	if low == 0 {
		ll = 1
	}
	for t > 0 {
		ll, t = ll+1, t/10
	}
	t = high
	for t > 0 {
		lh, t = lh+1, t/10
	}
	if ll == 1 {
		o = append(o, 0)
	}

	var cal func(int, int, int)
	cal = func(p, j, s int) {
		if p == 1 {
			if t := s + j; t >= low && t <= high {
				o = append(o, t)
			}
			return
		}
		if j > 0 {
			cal(p/10, j-1, s+p*j)
		}

		if j < 9 {
			cal(p/10, j+1, s+p*j)
		}
	}
	for i := ll; i <= lh; i++ {
		p := 1
		for j := 1; j < i; j++ {
			p *= 10
		}
		for j := 1; j <= 9; j++ {
			cal(p, j, 0)
		}
	}
	return o
}
