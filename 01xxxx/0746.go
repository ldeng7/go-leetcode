func minCostClimbingStairs(cost []int) int {
	l := len(cost)
	for i := 2; i < len(cost); i++ {
		if cost[i-1] < cost[i-2] {
			cost[i] += cost[i-1]
		} else {
			cost[i] += cost[i-2]
		}
	}
	if cost[l-1] < cost[l-2] {
		return cost[l-1]
	}
	return cost[l-2]
}
