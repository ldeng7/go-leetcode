import "sort"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	o, l := 0, len(arr2)
	sort.Ints(arr2)
	for _, a := range arr1 {
		i := sort.Search(l, func(j int) bool {
			return arr2[j] >= a
		})
		if (i == l || arr2[i]-a > d) && (i == 0 || a-arr2[i-1] > d) {
			o++
		}
	}
	return o
}
