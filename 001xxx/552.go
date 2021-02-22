import "sort"

func maxDistance(position []int, m int) int {
	sort.Ints(position)
	n := len(position)
	o, l, r := -1, 1, position[n-1]-position[0]
	for l <= r {
		h, p, c := l+(r-l)>>1, position[0], 1
		for i := 1; i < n; i++ {
			if v := position[i]; v-p >= h {
				p, c = v, c+1
			}
		}
		if c >= m {
			o, l = h, h+1
		} else {
			r = h - 1
		}
	}
	return o
}
