func maxNonOverlapping(nums []int, target int) int {
	o, s, p := 0, 0, 0
	m := map[int]int{0: 0}
	for i, n := range nums {
		s += n
		if j, ok := m[s-target]; ok && j >= p {
			o, p = o+1, i+1
		}
		m[s] = i + 1
	}
	return o
}
