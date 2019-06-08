func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minCost(costs [][]int) int {
	if 0 == len(costs) {
		return 0
	}
	for i := 1; i < len(costs); i++ {
		costs[i][0] += min(costs[i-1][1], costs[i-1][2])
		costs[i][1] += min(costs[i-1][0], costs[i-1][2])
		costs[i][2] += min(costs[i-1][0], costs[i-1][1])
	}
	c := costs[len(costs)-1]
	return min(min(c[0], c[1]), c[2])
}
