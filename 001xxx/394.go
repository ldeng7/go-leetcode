func findLucky(arr []int) int {
	m := map[int]int{}
	for _, a := range arr {
		m[a]++
	}
	o := -1
	for a, c := range m {
		if a == c && a > o {
			o = a
		}
	}
	return o
}
