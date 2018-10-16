import "sort"

func isPerfectSquare(num int) bool {
	if num <= 1 {
		return true
	}
	n := sort.Search(num, func(i int) bool {
		return i*i >= num
	})
	return n*n == num
}
