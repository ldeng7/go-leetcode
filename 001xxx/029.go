import "sort"

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][1]-costs[i][0] > costs[j][1]-costs[j][0]
	})
	sum := 0
	for _, c := range costs[:len(costs)>>1] {
		sum += c[0]
	}
	for _, c := range costs[len(costs)>>1:] {
		sum += c[1]
	}
	return sum
}
