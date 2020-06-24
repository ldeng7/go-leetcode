func buildArray(target []int, n int) []string {
	l := len(target)
	o := make([]string, 0, target[l-1]*2-l)
	m := 1
	for _, n := range target {
		for i := m; i < n; i++ {
			o = append(o, "Push", "Pop")
		}
		m = n + 1
		o = append(o, "Push")
	}
	return o
}
