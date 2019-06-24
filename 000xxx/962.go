import "sort"

func maxWidthRamp(A []int) int {
	l := len(A)
	ps := make([][2]int, l)
	for i, a := range A {
		ps[i][0], ps[i][1] = i, a
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i][1] < ps[j][1] || (ps[i][1] == ps[j][1] && ps[i][0] < ps[j][0])
	})
	o, m := 0, l
	for i := 0; i < l; i++ {
		j := ps[i][0]
		if j < m {
			m = j
		} else if j-m > o {
			o = j - m
		}
	}
	return o
}
