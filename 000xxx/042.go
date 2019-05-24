func trap(height []int) int {
	out := 0
	st := []int{}
	for i, h := range height {
		j := len(st) - 1
		for -1 != j && h > height[st[j]] {
			t := st[j]
			st = st[:j]
			j--
			if -1 == j {
				break
			}
			t1 := st[j]
			if h < height[t1] {
				out += (h - height[t]) * (i - t1 - 1)
			} else {
				out += (height[t1] - height[t]) * (i - t1 - 1)
			}
		}
		st = append(st, i)
	}
	return out
}
