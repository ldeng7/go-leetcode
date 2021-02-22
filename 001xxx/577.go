func cal(m1, m2 map[int]int) int {
	o := 0
	for n1, c1 := range m1 {
		s := n1 * n1
		for n2, c2 := range m2 {
			if s%n2 != 0 {
				continue
			}
			n3 := s / n2
			if n2 == n3 {
				o += c1 * c2 * (c2 - 1) / 2
			} else if c3, ok := m2[n3]; ok && n2 < n3 {
				o += c1 * c2 * c3
			}
		}
	}
	return o
}
func numTriplets(nums1 []int, nums2 []int) int {
	m1, m2 := make(map[int]int, len(nums1)), make(map[int]int, len(nums2))
	for _, n := range nums1 {
		m1[n]++
	}
	for _, n := range nums2 {
		m2[n]++
	}
	return cal(m1, m2) + cal(m2, m1)
}
