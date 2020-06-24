func numTeams(rating []int) int {
	o, l := 0, len(rating)
	if l < 3 {
		return 0
	}
	for i := 1; i < l-1; i++ {
		var l0, l1, g0, g1 int
		ri, j := rating[i], 0
		for ; j <= i; j++ {
			if rj := rating[j]; rj < ri {
				l0++
			} else if rj > ri {
				g0++
			}
		}
		for ; j < l; j++ {
			if rj := rating[j]; rj < ri {
				l1++
			} else if rj > ri {
				g1++
			}
		}
		o += l0*g1 + l1*g0
	}
	return o
}
