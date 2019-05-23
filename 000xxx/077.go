func combine(n int, k int) [][]int {
	if k < 0 || k > n {
		return [][]int{}
	} else if k == 0 {
		return [][]int{{}}
	}
	out := combine(n-1, k-1)
	for i, c := range out {
		out[i] = append(c, n)
	}
	for _, c := range combine(n-1, k) {
		out = append(out, c)
	}
	return out
}
