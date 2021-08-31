func checkPowersOfThree(n int) bool {
	for n != 0 {
		if r := n % 3; r == 0 || r == 1 {
			n = n / 3
		} else if r == 2 {
			return false
		}
	}
	return true
}
