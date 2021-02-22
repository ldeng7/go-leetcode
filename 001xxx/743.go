func restoreArray(adjacentPairs [][]int) []int {
	t := map[int]*[2]int{}
	for _, p := range adjacentPairs {
		a, b := p[0], p[1]
		pa, pb := t[a], t[b]
		if pa != nil {
			(*pa)[1] = b
		} else {
			p := [2]int{b, a}
			t[a] = &p
		}
		if pb != nil {
			(*pb)[1] = a
		} else {
			p := [2]int{a, b}
			t[b] = &p
		}
	}

	l, j := len(adjacentPairs)+1, 0
	o := make([]int, l)
	for k, p := range t {
		if (*p)[1] == k {
			o[0] = k
			j = (*p)[0]
			break
		}
	}
	for i := 1; i < l; i++ {
		o[i] = j
		p := t[j]
		if a := (*p)[0]; a == o[i-1] {
			j = (*p)[1]
		} else {
			j = a
		}
	}
	return o
}
