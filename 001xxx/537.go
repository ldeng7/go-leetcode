func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxSum(nums1 []int, nums2 []int) int {
	o, s1, s2, i1, i2, l1, l2 := 0, 0, 0, 0, 0, len(nums1), len(nums2)
	for i1 < l1 && i2 < l2 {
		for i1 < l1 || i2 < l2 {
			if i1 == l1 {
				s2 += nums2[i2]
				i2++
			} else if i2 == l2 {
				s1 += nums1[i1]
				i1++
			} else if nums1[i1] < nums2[i2] {
				s1 += nums1[i1]
				i1++
			} else if nums1[i1] > nums2[i2] {
				s2 += nums2[i2]
				i2++
			} else {
				break
			}
		}
		if i1 < l1 && i2 < l2 {
			o += max(s1, s2) + nums1[i1]
			s1, s2, i1, i2 = 0, 0, i1+1, i2+1
		}
	}

	for ; i1 < l1; i1++ {
		s1 += nums1[i1]
	}
	for ; i2 < l2; i2++ {
		s2 += nums2[i2]
	}
	o += max(s1, s2)
	return o % 1000000007
}
