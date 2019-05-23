func C2n(n int) int {
	out := n
	for i := n - 1; i > n-2; i-- {
		out *= i
	}
	return out >> 1
}

func numPairsDivisibleBy60(time []int) int {
	m0, m1 := map[int]int{}, map[int]int{}
	c0, c30 := 0, 0
	for _, t := range time {
		t %= 60
		if t > 30 {
			if _, ok := m1[t]; !ok {
				m1[t] = 1
			} else {
				m1[t]++
			}
		} else if t == 30 {
			c30++
		} else if t > 0 {
			if _, ok := m0[t]; !ok {
				m0[t] = 1
			} else {
				m0[t]++
			}
		} else {
			c0++
		}
	}

	cnt := C2n(c0) + C2n(c30)
	for t, c := range m0 {
		if c1, ok := m1[60-t]; ok {
			cnt += c * c1
		}
	}
	return cnt
}
