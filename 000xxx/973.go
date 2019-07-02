import "sort"

func kClosest(points [][]int, K int) [][]int {
	l := len(points)
	ds, ds0 := make([]int, l), make([]int, l)
	for i, p := range points {
		ds[i] = p[0]*p[0] + p[1]*p[1]
	}
	copy(ds0, ds)
	sort.Ints(ds)
	o, de, j := make([][]int, K), ds[K-1], 0
	for i := 0; i < l; i++ {
		if ds0[i] <= de {
			o[j], j = points[i], j+1
		}
	}
	return o
}
