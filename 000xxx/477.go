func totalHammingDistance(nums []int) int {
	out := 0
	for i := uint(0); i < 32; i++ {
		c, b := 0, 1<<i
		for _, num := range nums {
			if num&b != 0 {
				c++
			}
		}
		out += c * (len(nums) - c)
	}
	return out
}
