import "math"

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func minDistance(height int, width int, tree []int, squirrel []int, nuts [][]int) int {
	out, d := 0, math.MinInt64
	for _, nut := range nuts {
		dist := abs(tree[0]-nut[0]) + abs(tree[1]-nut[1])
		out += dist * 2
		d = max(d, dist-abs(squirrel[0]-nut[0])-abs(squirrel[1]-nut[1]))
	}
	return out - d
}
