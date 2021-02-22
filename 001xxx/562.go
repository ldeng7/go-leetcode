func findLatestStep(arr []int, m int) int {
	l := len(arr)
	if l == m {
		return m
	}
	t := make([]int, l+2)
	o, c := -1, 0
	for _, a := range arr {
		l, r := t[a-1], t[a+1]
		if l == m || r == m {
			o = c
		}
		t[a-l], t[a+r] = l+r+1, l+r+1
		c++
	}
	return o
}
