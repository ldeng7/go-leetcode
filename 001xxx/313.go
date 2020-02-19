func decompressRLElist(nums []int) []int {
	o := make([]int, 0, 32)
	for i := 0; i < len(nums); i += 2 {
		a, b := nums[i], nums[i+1]
		for j := 0; j < a; j++ {
			o = append(o, b)
		}
	}
	return o
}
