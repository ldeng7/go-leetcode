import "sort"

func min(ar, b int) int {
	if ar <= b {
		return ar
	}
	return b
}

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	o := 0
	for _, a := range arr {
		o = min(o+1, a)
	}
	return o
}
