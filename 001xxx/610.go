import (
	"math"
	"sort"
)

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func visiblePoints(points [][]int, angle int, location []int) int {
	c, l0, l1 := 0, location[0], location[1]
	ar := make([]float64, 0, len(points)*2)
	for _, p := range points {
		x, y := p[0]-l0, p[1]-l1
		if x == 0 && y == 0 {
			c++
		} else {
			ar = append(ar, math.Atan2(float64(y), float64(x))*180/math.Pi)
		}
	}
	sort.Float64s(ar)
	o, l := 0, len(ar)
	for i := 0; i < l; i++ {
		ar = append(ar, ar[i]+360)
	}
	l <<= 1
	for i, j := 0, 0; i < l; i++ {
		for ; j+1 < l && ar[j+1]-ar[i] <= float64(angle); j++ {
		}
		o = max(o, j-i+1)
	}
	return o + c
}
