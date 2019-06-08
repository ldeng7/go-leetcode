func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxDistance(arrays [][]int) int {
	a0, a1 := arrays[0], arrays[1]
	il, ir, l1, l2, r1, r2 := 0, 0, a0[0], a1[0], a0[len(a0)-1], a1[len(a1)-1]
	if l2 < l1 {
		l1, l2, il = l2, l1, 1
	}
	if r2 > r1 {
		r1, r2, ir = r2, r1, 1
	}
	for i := 2; i < len(arrays); i++ {
		a := arrays[i]
		l, r := a[0], a[len(a)-1]
		if l < l2 {
			l2 = l
			if l2 < l1 {
				l1, l2, il = l2, l1, i
			}
		}
		if r > r2 {
			r2 = r
			if r2 > r1 {
				r1, r2, ir = r2, r1, i
			}
		}
	}
	if il != ir {
		return r1 - l1
	}
	return max(r1-l2, r2-l1)
}
