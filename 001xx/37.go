func singleNumber(nums []int) int {
	h, l := 0, 0
	for _, n := range nums {
		l = (l ^ n) & (^h)
		h = (h ^ n) & (^l)
	}
	return l
}
