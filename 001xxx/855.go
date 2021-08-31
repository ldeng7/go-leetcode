func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxDistance(nums1 []int, nums2 []int) int {
	o, l1, l2 := 0, len(nums1), len(nums2)
	for i, j := 0, 0; i < l1 && j < l2; i++ {
		for ; j < l2 && nums2[j] >= nums1[i]; j++ {
		}
		if j > i {
			o = max(o, j-i-1)
		}
	}
	return o
}
