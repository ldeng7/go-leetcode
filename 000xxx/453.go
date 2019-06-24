import "math"

func minMoves(nums []int) int {
	s, m := 0, math.MaxInt64
	for _, n := range nums {
		s += n
		if n < m {
			m = n
		}
	}
	return s - m*len(nums)
}
