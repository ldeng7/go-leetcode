func dominantIndex(nums []int) int {
	f, s := -1, -1
	out := -1
	for i, n := range nums {
		if n > f {
			f, s = n, f
			out = i
		} else if n > s {
			s = n
		}
	}
	if f < s*2 {
		return -1
	}
	return out
}
