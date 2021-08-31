import "math"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minimumSize(nums []int, maxOperations int) int {
	l, r := 0, 0
	for _, n := range nums {
		r = max(r, n)
	}
	for l < r {
		m := l + (r-l)>>1
		s := 0
		for _, n := range nums {
			s += int(math.Ceil(float64(n)/float64(m))) - 1
		}
		if s > maxOperations {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
