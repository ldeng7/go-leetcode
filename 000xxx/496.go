func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for i, n := range nums2 {
		m[n] = i
	}
	out := make([]int, len(nums1))
	for i, n := range nums1 {
		j := m[n]
		for ; j < len(nums2); j++ {
			if nums2[j] > n {
				out[i] = nums2[j]
				break
			}
		}
		if j == len(nums2) {
			out[i] = -1
		}
	}
	return out
}
