func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	mc, dt := maxChoosableInteger, desiredTotal
	if mc >= dt {
		return true
	} else if mc*(mc+1)>>1 < dt {
		return false
	}
	m := map[int]bool{}

	var cal func(int, int) bool
	cal = func(t, u int) bool {
		if b, ok := m[u]; ok {
			return b
		}
		for i := 0; i < mc; i++ {
			j := (1 << uint(i))
			if (j&u) == 0 && (t <= i+1 || !cal(t-i-1, j|u)) {
				m[u] = true
				return true
			}
		}
		m[u] = false
		return false
	}
	return cal(dt, 0)
}
