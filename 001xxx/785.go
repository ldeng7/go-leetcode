func minElements(nums []int, limit int, goal int) int {
	s := -goal
	for _, n := range nums {
		s += n
	}
	if s < 0 {
		s = -s
	}
	return (s + limit - 1) / limit
}
