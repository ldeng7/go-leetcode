func singleNumber(nums []int) []int {
	d := 0
	for _, n := range nums {
		d ^= n
	}
	d &= -d
	a, b := 0, 0
	for _, n := range nums {
		if n&d != 0 {
			a ^= n
		} else {
			b ^= n
		}
	}
	return []int{a, b}
}
