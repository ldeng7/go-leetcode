func findSpecialInteger(arr []int) int {
	l := len(arr)
	l1 := l >> 2
	j, p := 0, arr[0]
	for i := 1; i < l; i++ {
		n := arr[i]
		if n != p {
			if i-j > l1 {
				return p
			}
			j, p = i, n
		}
	}
	return arr[l-1]
}
