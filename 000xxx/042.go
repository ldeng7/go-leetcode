func trap(height []int) int {
	o := 0
	st := []int{}
	for i, h := range height {
		j := len(st) - 1
		for -1 != j && h > height[st[j]] {
			t := st[j]
			st, j = st[:j], j-1
			if -1 == j {
				break
			}
			t1 := st[j]
			if h < height[t1] {
				o += (h - height[t]) * (i - t1 - 1)
			} else {
				o += (height[t1] - height[t]) * (i - t1 - 1)
			}
		}
		st = append(st, i)
	}
	return o
}
