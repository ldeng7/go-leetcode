import "sort"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxHeight(cuboids [][]int) int {
	for _, c := range cuboids {
		sort.Ints(c)
	}
	sort.Slice(cuboids, func(i, j int) bool {
		a, b := cuboids[i], cuboids[j]
		if a[0] != b[0] {
			return a[0] > b[0]
		} else if a[1] != b[1] {
			return a[1] > b[1]
		}
		return a[2] > b[2]
	})
	o, l := 0, len(cuboids)
	t := make([]int, l)
	for i := 0; i < l; i++ {
		ci := cuboids[i]
		for j := i - 1; j >= 0; j-- {
			if cuboids[j][1] >= ci[1] && cuboids[j][2] >= ci[2] {
				t[i] = max(t[j], t[i])
			}
		}
		t[i] += ci[2]
		o = max(o, t[i])
	}
	return o
}
