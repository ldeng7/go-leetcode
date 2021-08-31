func minOperations(nums1 []int, nums2 []int) int {
	s1, s2, m, n := 0, 0, len(nums1), len(nums2)
	cs1, cs2 := make([]int, 7), make([]int, 7)
	for _, num := range nums1 {
		s1 += num
		cs1[num]++
	}
	for _, num := range nums2 {
		s2 += num
		cs2[num]++
	}
	if s1 == s2 {
		return 0
	} else if 6*m < n || 6*n < m {
		return -1
	}
	if s2 < s1 {
		s1, s2 = s2, s1
		cs1, cs2 = cs2, cs1
	}

	o := 0
	for i := 1; i <= 6; i++ {
		if s1+(6-i)*cs1[i] >= s2 {
			return o + (s2-s1-1)/(6-i) + 1
		} else {
			s1 += (6 - i) * cs1[i]
			o += cs1[i]
		}
		if s2-(6-i)*cs2[7-i] <= s1 {
			return o + (s2-s1-1)/(6-i) + 1
		} else {
			s2 -= (6 - i) * cs2[7-i]
			o += cs2[7-i]
		}
	}
	return -1
}
