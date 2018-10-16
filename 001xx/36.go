func singleNumber(nums []int) int {
	out := 0
	for _, n := range nums {
		out ^= n
	}
	return out
}
