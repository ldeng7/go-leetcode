func isIdealPermutation(arr []int) bool {
	for i, n := range arr {
		if n-i > 1 || i-n > 1 {
			return false
		}
	}
	return true
}
