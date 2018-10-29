func partitionLabels(s string) []int {
	out := []int{}
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]] = i
	}
	j, k := 0, 0
	for i := 0; i < len(s); i++ {
		if m[s[i]] > j {
			j = m[s[i]]
		}
		if i == j {
			out = append(out, i-k+1)
			k = i + 1
		}
	}
	return out
}
