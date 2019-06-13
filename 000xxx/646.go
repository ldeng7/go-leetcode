import (
	"math"
	"sort"
)

func findLongestChain(pairs [][]int) int {
	o, e := 0, math.MinInt64
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	for _, p := range pairs {
		if p[0] > e {
			o, e = o+1, p[1]
		}
	}
	return o
}
