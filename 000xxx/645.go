func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func findErrorNums(nums []int) []int {
	o := []int{-1, -1}
	for _, n := range nums {
		j := abs(n) - 1
		if nums[j] < 0 {
			o[0] = j + 1
		} else {
			nums[j] = -nums[j]
		}
	}
	for i, n := range nums {
		if n > 0 {
			o[1] = i + 1
		}
	}
	return o
}
