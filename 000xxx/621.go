func leastInterval(tasks []byte, n int) int {
	m, c := 0, 0
	cs := [26]int{}
	for _, t := range tasks {
		t -= 'A'
		cs[t]++
		ct := cs[t]
		if m == ct {
			c++
		} else if m < ct {
			m, c = ct, 1
		}
	}
	nIdles := (m-1)*(n-c+1) - (len(tasks) - m*c)
	if nIdles < 0 {
		nIdles = 0
	}
	return len(tasks) + nIdles
}
