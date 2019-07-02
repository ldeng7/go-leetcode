func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	ds := make([]int, 0, 8)
	for x > 0 {
		ds = append(ds, x%10)
		x /= 10
	}
	l := len(ds)
	for i := 0; i < l>>1; i++ {
		if ds[i] != ds[l-1-i] {
			return false
		}
	}
	return true
}
