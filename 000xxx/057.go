func insert(intervals [][]int, newInterval []int) [][]int {
	out := [][]int{}
	i := 0
	for _, interval := range intervals {
		if interval[1] < newInterval[0] {
			out = append(out, interval)
			i++
		} else if interval[0] > newInterval[1] {
			out = append(out, interval)
		} else {
			if interval[0] < newInterval[0] {
				newInterval[0] = interval[0]
			}
			if interval[1] > newInterval[1] {
				newInterval[1] = interval[1]
			}
		}
	}
	out = append(out, nil)
	copy(out[i+1:], out[i:len(out)-1])
	out[i] = newInterval
	return out
}
