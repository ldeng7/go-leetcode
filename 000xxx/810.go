func xorGame(nums []int) bool {
	n, l := 0, len(nums)
	for _, num := range nums {
		n ^= num
	}
	return n == 0 || l&1 == 0
}
