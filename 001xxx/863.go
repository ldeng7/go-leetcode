func subsetXORSum(nums []int) int {
	t := 0
	for _, n := range nums {
		t |= n
	}
	return t << (len(nums) - 1)
}
