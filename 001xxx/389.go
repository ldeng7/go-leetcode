func createTargetArray(nums []int, index []int) []int {
	l := len(nums)
	o := make([]int, 0, l)
	for i, j := range index {
		if j != len(o) {
			o = append(o, 0)
			copy(o[j+1:], o[j:])
			o[j] = nums[i]
		} else {
			o = append(o, nums[i])
		}
	}
	return o
}
