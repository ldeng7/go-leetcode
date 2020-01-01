func smallestDivisor(nums []int, threshold int) int {
	o, l, r := -1, 1, nums[0]
	for _, n := range nums {
		if n > r {
			r = n
		}
	}
	for l <= r {
		s, m := 0, l+(r-l)>>1
		for _, n := range nums {
			s += (n-1)/m + 1
		}
		if s <= threshold {
			o, r = m, m-1
		} else {
			l = m + 1
		}
	}
	return o
}
