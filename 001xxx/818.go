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

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	s, md, d := 0, 0, 10001
	for i, n1 := range nums1 {
		n2 := nums2[i]
		t := abs(n1 - n2)
		s += t
		if md < t {
			for _, n1 := range nums1 {
				d = min(d, abs(n1-n2))
			}
			md, d = max(md, t-d), 10005
		}
	}
	return (s - md) % 1000000007
}
