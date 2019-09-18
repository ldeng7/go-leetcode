func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	ar, i, j, l1, l2 := []int{}, 0, 0, len(nums1), len(nums2)
	for i < l1 && j < l2 {
		n1, n2 := nums1[i], nums2[j]
		if n1 < n2 {
			ar, i = append(ar, n1), i+1
		} else if n1 > n2 {
			ar, j = append(ar, n2), j+1
		} else {
			ar, i, j = append(ar, n1, n2), i+1, j+1
		}
	}
	if i < l1 {
		ar = append(ar, nums1[i:]...)
	}
	if j < l2 {
		ar = append(ar, nums2[j:]...)
	}
	l := len(ar)
	if l&1 == 0 {
		return float64(ar[l/2-1]+ar[l/2]) / 2.0
	}
	return float64(ar[l/2])
}
