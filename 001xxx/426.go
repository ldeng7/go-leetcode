func countElements(arr []int) int {
	m := map[int]bool{}
	for _, a := range arr {
		m[a] = true
	}
	o := 0
	for _, a := range arr {
		if m[a+1] {
			o++
		}
	}
	return o
}
