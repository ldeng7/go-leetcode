import "sort"

func outerTrees(points [][]int) [][]int {
	l := len(points)
	sort.Slice(points, func(i, j int) bool {
		a, b := points[i], points[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	o, k := make([][]int, l<<1), 0

	externalProd := func(i, j, k int) int {
		a, b, c := o[i], o[j], points[k]
		return (b[0]-a[0])*(c[1]-b[1]) - (b[1]-a[1])*(c[0]-b[0])
	}
	for i := 0; i < l; i++ {
		for k > 1 && externalProd(k-2, k-1, i) < 0 {
			k--
		}
		o[k], k = points[i], k+1
	}
	for i, k0 := l-2, k; i >= 0; i-- {
		for k > k0 && externalProd(k-2, k-1, i) < 0 {
			k--
		}
		o[k], k = points[i], k+1
	}

	o = o[:k]
	sort.Slice(o, func(i, j int) bool {
		a, b := o[i], o[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	n := 0
	for i := 1; i < k; i++ {
		a, b := o[i], o[i-n]
		if a[0] == b[0] && a[1] == b[1] {
			n++
		} else {
			o[i-n] = a
		}
	}
	return o[:k-n]
}
