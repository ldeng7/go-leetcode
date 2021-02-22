func threeConsecutiveOdds(arr []int) bool {
	c := 0
	for _, a := range arr {
		if a&1 == 1 {
			c++
			if c == 3 {
				return true
			}
		} else {
			c = 0
		}
	}
	return false
}
