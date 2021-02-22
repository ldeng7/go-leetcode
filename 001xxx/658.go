func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minOperations(nums []int, x int) int {
	t := -x
	for _, n := range nums {
		t += n
	}
	o, n := 100001, len(nums)
	for l, r, s := 0, 0, 0; r < n; r++ {
		s += nums[r]
		for s > t && l <= r {
			s -= nums[l]
			l++
		}
		if s == t {
			o = min(o, n-(r-l+1))
		}
	}
	if o == 100001 {
		return -1
	}
	return o
}
