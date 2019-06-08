func findMaxAverage(nums []int, k int) float64 {
	sums := make([]float64, len(nums)+1)
	l, r := float64(nums[0]), float64(nums[0])
	for _, n := range nums {
		fn := float64(n)
		if fn < l {
			l = fn
		}
		if fn > r {
			r = fn
		}
	}
	for r-l > 1e-5 {
		var min, m float64 = 0.0, l + (r-l)>>1
		b := false
		for i := 1; i <= len(nums); i++ {
			sums[i] = sums[i-1] + float64(nums[i-1]) - m
			if i >= k {
				if sums[i-k] < min {
					min = sums[i-k]
				}
				if sums[i] > min {
					b = true
					break
				}
			}
		}
		if b {
			l = m
		} else {
			r = m
		}
	}
	return l
}
