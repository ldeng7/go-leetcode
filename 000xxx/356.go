import "math"

func isReflected(points [][]int) bool {
	m := map[int]map[int]bool{}
	max, min := math.MinInt64, math.MaxInt64
	for _, p := range points {
		if p[0] > max {
			max = p[0]
		}
		if p[0] < min {
			min = p[0]
		}
		if mi, ok := m[p[0]]; ok {
			mi[p[1]] = true
		} else {
			m[p[0]] = map[int]bool{p[1]: true}
		}
	}
	for _, p := range points {
		i := max + min - p[0]
		if _, ok := m[i]; !ok {
			return false
		}
		if _, ok := m[i][p[1]]; !ok {
			return false
		}
	}
	return true
}
