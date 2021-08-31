import "sort"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	o, t := 1, k
	for l, r := 0, 0; r < len(nums)-1; {
		if t >= 0 {
			r++
			t -= (nums[r] - nums[r-1]) * (r - l)
		} else {
			t += nums[r] - nums[l]
			l++
		}
		if t >= 0 {
			o = max(o, r-l+1)
		}
	}
	return o
}
