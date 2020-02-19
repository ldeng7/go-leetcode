func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func dis(a, b int) int {
	return abs(a/6-b/6) + abs(a%6-b%6)
}

func minimumDistance(word string) int {
	o, l, m := 0, len(word), 0
	t := [26]int{}
	for i := 0; i < l-1; i++ {
		a, b := int(word[i]-'A'), int(word[i+1]-'A')
		d := dis(a, b)
		for j := 0; j < 26; j++ {
			t[a] = max(t[a], t[j]+d-dis(j, b))
		}
		m = max(m, t[a])
		o += d
	}
	return o - m
}
