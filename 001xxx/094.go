func carPooling(trips [][]int, capacity int) bool {
	t := [1001]int{}
	for _, trip := range trips {
		t[trip[1]] += trip[0]
		t[trip[2]] -= trip[0]
	}
	for _, n := range t {
		if 0 == n {
			continue
		}
		capacity -= n
		if capacity < 0 {
			return false
		}
	}
	return true
}
