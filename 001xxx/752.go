func check(nums []int) bool {
	l := len(nums)
	for i := 0; i < l; i++ {
		j := 1
		for ; j < l; j++ {
			if nums[(i+j)%l] < nums[(i+j-1)%l] {
				break
			}
		}
		if j == l {
			return true
		}
	}
	return false
}
