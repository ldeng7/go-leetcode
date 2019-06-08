func leastBricks(wall [][]int) int {
	max, m := 0, map[int]int{}
	for _, line := range wall {
		s := 0
		for i := 0; i < len(line)-1; i++ {
			s += line[i]
			m[s]++
			if m[s] > max {
				max = m[s]
			}
		}
	}
	return len(wall) - max
}
