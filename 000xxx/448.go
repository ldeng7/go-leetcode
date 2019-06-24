func findDisappearedNumbers(nums []int) []int {
	o, l := []int{}, len(nums)
	for i := 0; i < l; i++ {
		nums[(nums[i]-1)%l] += l
	}
	for i := 0; i < l; i++ {
		if nums[i] <= l {
			o = append(o, i+1)
		}
	}
	return o
}
