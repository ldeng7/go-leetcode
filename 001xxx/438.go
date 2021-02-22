func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func longestSubarray(nums []int, limit int) int {
	ma, mi := make([]int, 0, 16), make([]int, 0, 16)
	o, l := 0, len(nums)
	for i, j := 0, 0; i < l; i++ {
		n := nums[i]
		for len(ma) > 0 && n > ma[len(ma)-1] {
			ma = ma[:len(ma)-1]
		}
		for len(mi) > 0 && n < mi[len(mi)-1] {
			mi = mi[:len(mi)-1]
		}
		ma = append(ma, n)
		mi = append(mi, n)
		for ; j < l && ma[0]-mi[0] > limit; j++ {
			if nums[j] == ma[0] {
				ma = ma[1:]
			}
			if nums[j] == mi[0] {
				mi = mi[1:]
			}
		}
		o = max(o, i-j+1)
	}
	return o
}
