import "math"

func thirdMax(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
	for _, n := range nums {
		if n > a {
			a, b, c = n, a, b
		} else if n < a && n > b {
			b, c = n, b
		} else if n < b && n > c {
			c = n
		}
	}
	if c == math.MinInt64 {
		return a
	}
	return c
}
