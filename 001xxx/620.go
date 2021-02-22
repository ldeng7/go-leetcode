import "math"

func bestCoordinate(towers [][]int, radius int) []int {
	q := make([][2]int, len(towers), len(towers)+32)
	for i, s1 := range towers {
		q[i] = [2]int{s1[0], s1[1]}
	}
	var o [2]int
	var s, r float64 = 0, float64(radius)
	for 0 != len(q) {
		x, y := q[0][0], q[0][1]
		q = q[1:]
		var s1 float64
		for _, t := range towers {
			tx, ty := t[0], t[1]
			d := math.Sqrt(float64((tx-x)*(tx-x) + (ty-y)*(ty-y)))
			if d > r {
				continue
			}
			s1 += math.Floor(float64(t[2]) / (d + 1))
		}

		if s1 > s {
			s = s1
			o = [2]int{x, y}
			q = append(q, [2]int{x - 1, y - 1}, [2]int{x - 1, y}, [2]int{x - 1, y + 1},
				[2]int{x, y - 1}, [2]int{x, y + 1}, [2]int{x + 1, y - 1}, [2]int{x + 1, y}, [2]int{x + 1, y + 1})
		} else if s1 == s {
			if x < o[0] || (x == o[0] && y < o[1]) {
				o = [2]int{x, y}
			}
		}
	}
	return o[:]
}
