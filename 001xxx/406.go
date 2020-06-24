func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func stoneGameIII(stoneValue []int) string {
	l := len(stoneValue)
	stoneValue = append(stoneValue, 0, 0)
	t := make([]int, l+3)
	for i := l - 1; i >= 0; i-- {
		s0, s1 := stoneValue[i], stoneValue[i+1]
		t[i] = max(s0-t[i+1], max(s0+s1-t[i+2], s0+s1+stoneValue[i+2]-t[i+3]))
	}
	if t0 := t[0]; t0 > 0 {
		return "Alice"
	} else if t0 < 0 {
		return "Bob"
	} else {
		return "Tie"
	}
}
