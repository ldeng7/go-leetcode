func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func mincostTickets(days []int, costs []int) int {
	t, s := [366]int{}, map[int]bool{}
	for i := 0; i <= 365; i++ {
		t[i] = -1
	}
	for _, d := range days {
		s[d] = true
	}

	var cal func(int) int
	cal = func(i int) int {
		if i > 365 {
			return 0
		} else if t[i] >= 0 {
			return t[i]
		}
		o := 0
		if s[i] {
			o = min(cal(i+1)+costs[0], cal(i+7)+costs[1])
			o = min(o, cal(i+30)+costs[2])
		} else {
			o = cal(i + 1)
		}
		t[i] = o
		return o
	}
	return cal(1)
}
