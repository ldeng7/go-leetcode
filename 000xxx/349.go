func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]byte{}
	for _, n := range nums1 {
		m[n] = 1
	}
	for _, n := range nums2 {
		m[n] |= 2
	}
	out := []int{}
	for n, c := range m {
		if c == 3 {
			out = append(out, n)
		}
	}
	return out
}
