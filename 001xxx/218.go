func longestSubsequence(arr []int, difference int) int {
	o, m := 1, map[int]int{}
	for _, a := range arr {
		t := m[a-difference] + 1
		m[a] = t
		if t > o {
			o = t
		}
	}
	return o
}
