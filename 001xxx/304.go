func sumZero(n int) []int {
	m := n >> 1
	o, i := make([]int, n), 0
	for j := -m; j < 0; j++ {
		o[i] = j
		i++
	}
	if n&1 == 1 {
		o[i] = 0
		i++
	}
	for j := 1; j <= m; j++ {
		o[i] = j
		i++
	}
	return o
}
