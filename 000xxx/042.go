func trap(height []int) int {
	st := []int{}
	var i, out int
	for i < len(height) {
		if 0 == len(st) || height[i] <= height[st[len(st)-1]] {
			st = append(st, i)
			i++
		} else {
			t := st[len(st)-1]
			st = st[:len(st)-1]
			if 0 == len(st) {
				continue
			}
			t1 := st[len(st)-1]
			if height[i] < height[t1] {
				out += (height[i] - height[t]) * (i - t1 - 1)
			} else {
				out += (height[t1] - height[t]) * (i - t1 - 1)
			}
		}
	}
	return out
}
