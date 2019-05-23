func findMaxConsecutiveOnes(nums []int) int {
	c, m := 0, 0
	nums = append(nums, 0)
	for _, n := range nums {
		if n == 1 {
			c++
		} else {
			if c > m {
				m = c
			}
			c = 0
		}
	}
	return m
}
