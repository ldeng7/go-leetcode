import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

var patterns = [4][2]uint8{{0, 1}, {0, 3}, {2, 3}, {2, 1}}

func isRectangleCover(rectangles [][]int) bool {
	s := map[string]bool{}
	minY, minX, maxY, maxX, area := math.MaxInt64, math.MaxInt64, math.MinInt64, math.MinInt64, 0
	for _, rect := range rectangles {
		minY = min(minY, rect[0])
		minX = min(minX, rect[1])
		maxY = max(maxY, rect[2])
		maxX = max(maxX, rect[3])
		area += (rect[2] - rect[0]) * (rect[3] - rect[1])
		for i := 0; i < 4; i++ {
			k := fmt.Sprintf("%d:%d", rect[patterns[i][0]], rect[patterns[i][1]])
			s[k] = !s[k]
		}
	}
	k1 := fmt.Sprintf("%d:%d", minY, minX)
	k2 := fmt.Sprintf("%d:%d", minY, maxX)
	k3 := fmt.Sprintf("%d:%d", maxY, maxX)
	k4 := fmt.Sprintf("%d:%d", maxY, minX)
	if !s[k1] || !s[k2] || !s[k3] || !s[k4] {
		return false
	}
	c := 0
	for _, v := range s {
		if v {
			c++
		}
	}
	if 4 != c {
		return false
	}
	return area == (maxY-minY)*(maxX-minX)
}
