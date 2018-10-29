func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	ds := []int{}
	for x > 0 {
		ds = append(ds, x%10)
		x /= 10
	}
	for i := 0; i < (len(ds) >> 1); i++ {
		if ds[i] != ds[len(ds)-1-i] {
			return false
		}
	}
	return true
}
