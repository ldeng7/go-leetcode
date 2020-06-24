import "sort"

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func getStrongest(arr []int, k int) []int {
	sort.Ints(arr)
	l, r := 0, len(arr)-1
	m := arr[r/2]
	o := make([]int, k)
	for i := 0; i < k; i++ {
		a, b := arr[l], arr[r]
		if abs(m-a) <= abs(b-m) {
			o[i], r = b, r-1
		} else {
			o[i], l = a, l+1
		}
	}
	return o
}
