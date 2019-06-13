func findMaxLength(nums []int) int {
	o, s := 0, 0
	m := map[int]int{0: -1}
	for i, n := range nums {
		s += (n << 1) - 1
		if j, ok := m[s]; ok {
			if i-j > o {
				o = i - j
			}
		} else {
			m[s] = i
		}
	}
	return o
}
