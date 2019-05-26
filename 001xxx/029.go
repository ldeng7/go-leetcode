import "sort"

type Pairs [][]int

func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Less(i, j int) bool {
	return p[i][1]-p[i][0] > p[j][1]-p[j][0]
}

func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func twoCitySchedCost(costs [][]int) int {
	var cs Pairs = costs
	sort.Sort(cs)
	sum := 0
	for _, c := range cs[:len(cs)>>1] {
		sum += c[0]
	}
	for _, c := range cs[len(cs)>>1:] {
		sum += c[1]
	}
	return sum
}
