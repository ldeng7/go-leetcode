func findPairs(nums []int, k int) int {
	o, m := 0, map[int]int{}
	for _, n := range nums {
		m[n]++
	}
	for n, c := range m {
		if k == 0 && c > 1 {
			o++
		} else if _, ok := m[n+k]; ok && k > 0 {
			o++
		}
	}
	return o
}
