func waysToMakeFair(nums []int) int {
	o, l, so, se := 0, len(nums), 0, 0
	aro, are := make([]int, l), make([]int, l)
	for i := l - 1; i >= 0; i-- {
		aro[i], are[i] = se, so
		if i&1 == 1 {
			so += nums[i]
		} else {
			se += nums[i]
		}
	}
	so, se = 0, 0
	for i := 0; i < l; i++ {
		if so+aro[i] == se+are[i] {
			o++
		}
		if i&1 == 1 {
			so += nums[i]
		} else {
			se += nums[i]
		}
	}
	return o
}
