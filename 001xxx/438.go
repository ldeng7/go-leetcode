func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func longestSubarray(nums []int, limit int) int {
	o, l := 0, len(nums)
	if l == 1 {
		return 1
	}
	for i, n := range nums {
		nMin, nMax := n, n
		iMin, iMax := i, i
		for j := i + 1; j < l; j++ {
			m := nums[j]
			if m <= nMin {
				nMin, iMin = m, j
			}
			if m >= nMax {
				nMax, iMax = m, j
			}
			if nMax-nMin > limit {
				o, i = max(o, j-i), min(iMax, iMin)
				break
			}
			o = max(o, j-i+1)
			if o == l-i {
				return o
			}
		}
	}
	return o
}
