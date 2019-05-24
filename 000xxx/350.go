func intersect(nums1 []int, nums2 []int) []int {
	out := []int{}
	m := map[int]int{}
	for _, n := range nums1 {
		m[n]++
	}
	for _, n := range nums2 {
		if 0 != m[n] {
			out = append(out, n)
			m[n]--
		}
	}
	return out
}
