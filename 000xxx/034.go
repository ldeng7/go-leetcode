func searchRange(nums []int, target int) []int {
	var cal func(int, int) []int
	cal = func(b, e int) []int {
		if b >= e {
			return []int{-1, -1}
		}
		m := b + (e-b)>>1
		if nums[m] == target {
			l, r := m, m
			for l >= 0 && nums[l] == target {
				l--
			}
			l++
			for r < len(nums) && nums[r] == target {
				r++
			}
			r--
			return []int{l, r}
		} else if nums[m] > target {
			return cal(b, m)
		}
		return cal(m+1, e)
	}
	return cal(0, len(nums))
}
