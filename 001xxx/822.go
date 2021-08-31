func arraySign(nums []int) int {
	var b uint64 = 1
	for _, n := range nums {
		if n != 0 {
			b ^= uint64(n) >> 63
		} else {
			return 0
		}
	}
	return (int(b) << 1) - 1
}
