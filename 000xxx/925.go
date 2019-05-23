func isLongPressedName(name string, typed string) bool {
	i0, i1 := 0, 0
	for i0 != len(name) {
		c := name[i0]
		n0 := 1
		for i0++; i0 < len(name); i0++ {
			if name[i0] != c {
				break
			}
			n0++
		}
		n1 := 0
		for ; i1 < len(typed); i1++ {
			if typed[i1] != c {
				break
			}
			n1++
		}
		if n1 < n0 {
			return false
		}
	}
	return true
}
