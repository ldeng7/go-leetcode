import "sort"

func maxNumberOfApples(arr []int) int {
	sort.Ints(arr)
	o, s := 0, 0
	for _, a := range arr {
		o, s = o+1, s+a
		if s > 5000 {
			return o - 1
		}
	}
	return o
}
