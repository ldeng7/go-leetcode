import "math"

func increasingTriplet(nums []int) bool {
	m1, m2 := math.MaxInt64, math.MaxInt64
	for _, n := range nums {
		if m1 >= n {
			m1 = n
		} else if m2 >= n {
			m2 = n
		} else {
			return true
		}
	}
	return false
}
