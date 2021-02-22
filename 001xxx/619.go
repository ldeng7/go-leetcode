import "sort"

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	l := len(arr)
	r := l / 20
	var s float64 = 0
	for i := r; i < l-r; i++ {
		s += float64(arr[i])
	}
	return s / float64(l-r*2)
}
