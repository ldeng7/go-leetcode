func singleNumber(nums []int) int {
	o := 0
	for _, n := range nums {
		o ^= n
	}
	return o
}
