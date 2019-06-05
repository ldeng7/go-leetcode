func canCompleteCircuit(gas []int, cost []int) int {
	s, b, m := 0, 0, -1
	for i := len(gas) - 1; i >= 0; i-- {
		s += gas[i] - cost[i]
		if s > m {
			b, m = i, s
		}
	}
	if s < 0 {
		return -1
	}
	return b
}
