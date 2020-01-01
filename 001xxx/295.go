func findNumbers(nums []int) int {
	o := 0
	for _, n := range nums {
		c := 0
		for ; n != 0; n /= 10 {
			c++
		}
		if c&1 == 0 {
			o++
		}
	}
	return o
}
