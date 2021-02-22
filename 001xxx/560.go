func mostVisited(n int, rounds []int) []int {
	var o []int
	b, e := rounds[0], rounds[len(rounds)-1]
	if b <= e {
		o = make([]int, e-b+1)
		for i := b; i <= e; i++ {
			o[i-b] = i
		}
	} else {
		o = make([]int, e+n-b+1)
		for i := 1; i <= e; i++ {
			o[i-1] = i
		}
		for i := b; i <= n; i++ {
			o[e+i-b] = i
		}
	}
	return o
}
