func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func countBalls(lowLimit int, highLimit int) int {
	m := map[int]int{}
	for i := lowLimit; i <= highLimit; i++ {
		n := 0
		for j := i; j != 0; j /= 10 {
			n += j % 10
		}
		m[n]++
	}
	o := 0
	for _, c := range m {
		o = max(o, c)
	}
	return o
}
