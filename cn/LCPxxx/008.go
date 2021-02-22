func getTriggerTime(increase [][]int, requirements [][]int) []int {
	li := len(increase)
	t := make([][3]int, li+1)
	for i := 0; i < li; i++ {
		ti, ti1, inc := t[i], &t[i+1], increase[i]
		(*ti1)[0], (*ti1)[1], (*ti1)[2] = ti[0]+inc[0], ti[1]+inc[1], ti[2]+inc[2]
	}
	o := make([]int, len(requirements))
	for i, req := range requirements {
		l, r := 0, li
		r0, r1, r2 := req[0], req[1], req[2]
		for l < r {
			m := l + (r-l)>>1
			if tm := t[m]; tm[0] >= r0 && tm[1] >= r1 && tm[2] >= r2 {
				r = m
			} else {
				l = m + 1
			}
		}
		if tl := t[l]; tl[0] >= r0 && tl[1] >= r1 && tl[2] >= r2 {
			o[i] = l
		} else {
			o[i] = -1
		}
	}
	return o
}
