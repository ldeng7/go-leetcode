func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, len(nums1)-1
	for j >= 0 {
		if i >= 0 && nums1[i] >= nums2[j] {
			nums1[k], i = nums1[i], i-1
		} else {
			nums1[k], j = nums2[j], j-1
		}
		k--
	}
}
