func cal(nums1 []int, nums2 []int, k int) int {
	l1, l2 := len(nums1), len(nums2)
	if 0 == l1 {
		return nums2[k-1]
	}
	if 0 == l2 {
		return nums1[k-1]
	}
	if k == 1 {
		if nums1[0] < nums2[0] {
			return nums1[0]
		}
		return nums2[0]
	}
	i, j := k/2, k/2
	if l1 < i {
		i = l1
	}
	if l2 < j {
		j = l2
	}
	if nums1[i-1] > nums2[j-1] {
		return cal(nums1, nums2[j:], k-j)
	} else {
		return cal(nums1[i:], nums2, k-i)
	}
	return 0
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	return float64(cal(nums1, nums2, (l1+l2+1)/2)+cal(nums1, nums2, (l1+l2+2)/2)) / 2.0
}
