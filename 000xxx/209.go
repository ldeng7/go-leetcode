func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if 0 == n {
		return 0
	}
	l, r, sum, o := 0, 0, 0, n+1
	for r < n {
		for sum < s && r < n {
			sum += nums[r]
			r++
		}
		for sum >= s {
			if r-l < o {
				o = r - l
			}
			sum -= nums[l]
			l++
		}
	}
	if o == n+1 {
		return 0
	}
	return o
}
