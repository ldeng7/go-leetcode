func gcd(a, b int) int {
	if 0 == b {
		return a
	}
	return gcd(b, a%b)
}

func isGoodArray(nums []int) bool {
	g := nums[0]
	for i := 1; i < len(nums); i++ {
		if g = gcd(g, nums[i]); g == 1 {
			return true
		}
	}
	return g == 1
}
