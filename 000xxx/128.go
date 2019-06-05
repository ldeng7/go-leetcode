func longestConsecutive(nums []int) int {
	out := 0
	m := map[int]int{}
	for _, num := range nums {
		if _, ok := m[num]; ok {
			continue
		}
		l, r := m[num-1], m[num+1]
		s := l + r + 1
		m[num], m[num-l], m[num+r] = s, s, s
		if s > out {
			out = s
		}
	}
	return out
}
