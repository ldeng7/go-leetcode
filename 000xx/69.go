import "sort"

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	n := sort.Search(x, func(i int) bool {
		return i*i > x
	})
	return n - 1
}
