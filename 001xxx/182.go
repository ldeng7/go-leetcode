import "sort"

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func shortestDistanceColor(colors []int, queries [][]int) []int {
	m := map[int][]int{}
	for i, c := range colors {
		m[c] = append(m[c], i)
	}
	o := []int{}
	for _, q := range queries {
		t := q[0]
		ar, ok := m[q[1]]
		if !ok {
			o = append(o, -1)
			continue
		}
		l := len(ar)
		j := sort.Search(l, func(i int) bool {
			return ar[i] >= t
		})
		if j == 0 {
			j = ar[0]
		} else if j == l {
			j = ar[l-1]
		} else if ar[j]-t < t-ar[j-1] {
			j = ar[j]
		} else {
			j = ar[j-1]
		}
		o = append(o, abs(j-t))
	}
	return o
}
