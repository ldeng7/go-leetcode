func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
	o := make([][]int, 0, len(intervals))
	b, e := toBeRemoved[0], toBeRemoved[1]
	for _, itv := range intervals {
		b1, e1 := itv[0], itv[1]
		if e1 <= b || b1 >= e {
			o = append(o, itv)
		} else {
			if b1 < b {
				o = append(o, []int{b1, b})
			}
			if e1 > e {
				o = append(o, []int{e, e1})
			}
		}
	}
	return o
}
