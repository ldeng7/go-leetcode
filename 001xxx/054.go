import "sort"

func rearrangeBarcodes(barcodes []int) []int {
	l, m := len(barcodes), make(map[int]int)
	for _, bar := range barcodes {
		m[bar]++
	}
	ar := make([][2]int, len(m))
	i := 0
	for bar, c := range m {
		ar[i], i = [2]int{bar, c}, i+1
	}
	sort.Slice(ar, func(i, j int) bool {
		return ar[i][1] > ar[j][1]
	})
	o := make([]int, len(barcodes))
	b := 0
	for _, e := range ar {
		bar := e[0]
		for c := e[1]; c > 0; c-- {
			o[b], b = bar, b+2
			if b >= l {
				b = 1
			}
		}
	}
	return o
}
