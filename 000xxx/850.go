import (
	"math"
	"sort"
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

func rectangleArea(rectangles [][]int) int {
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i][0] < rectangles[j][0]
	})
	var o, x, y, lx, ly, h int
	for _, r := range rectangles {
		h = max(h, r[3])
	}
	for y < h {
		x, lx, ly = 0, 0, math.MaxInt64
		b := false
		for _, r := range rectangles {
			if r[1] > y {
				ly = min(ly, r[1])
			}
			if x < r[2] && y >= r[1] && y < r[3] {
				lx, ly = lx+r[2]-max(x, r[0]), min(ly, r[3])
				x, b = r[2], true
			}
		}
		if b {
			o, y = o+lx*(ly-y), ly
		} else {
			y = ly
		}
	}
	return o % 1000000007
}
