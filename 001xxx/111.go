func maxDepthAfterSplit(seq string) []int {
	n := len(seq)
	o, l, r := make([]int, n), 0, 0
	for i := 0; i < n; i++ {
		if seq[i] == '(' {
			o[i], l = l, l^1
		} else {
			o[i], r = r, r^1
		}
	}
	return o
}
