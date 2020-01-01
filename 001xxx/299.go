func replaceElements(arr []int) []int {
	l := len(arr)
	m := arr[l-1]
	t := m
	arr[l-1] = -1
	for i := l - 2; i >= 0; i-- {
		if t > m {
			m = t
		}
		arr[i], t = m, arr[i]
	}
	return arr
}
