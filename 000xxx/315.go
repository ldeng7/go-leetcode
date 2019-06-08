func countSmaller(nums []int) []int {
	l, t := len(nums), []int{}
	out := make([]int, l)
	for i := l - 1; i >= 0; i-- {
		l, r := 0, len(t)
		for l < r {
			m := l + (r-l)>>1
			if t[m] >= nums[i] {
				r = m
			} else {
				l = m + 1
			}
		}
		out[i] = r
		t = append(t, 0)
		if r != len(t)-1 {
			copy(t[r+1:], t[r:])
		}
		t[r] = nums[i]
	}
	return out
}
