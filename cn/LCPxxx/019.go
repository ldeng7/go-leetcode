func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minimumOperations(leaves string) int {
	v := 1
	if 'r' == leaves[0] {
		v = -1
	}
	o, m := 1, v
	l := len(leaves)
	for i := 1; i < l; i++ {
		if 'y' == leaves[i] {
			v++
		} else {
			v--
		}
		if i != l-1 {
			o = min(o, m-v)
		}
		m = min(m, v)
	}
	return o + (v+l)/2
}
