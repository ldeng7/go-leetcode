import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	o := [][]int{}
	m, l := math.MaxInt64, len(arr)
	for i := 0; i < l-1; i++ {
		n, nn := arr[i], arr[i+1]
		d := nn - n
		if m > d {
			o = [][]int{{n, nn}}
			m = d
		} else if m == d {
			o = append(o, []int{n, nn})
		}
	}
	return o
}
