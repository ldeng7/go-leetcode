func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func findDuplicates(nums []int) []int {
	o := []int{}
	for _, n := range nums {
		i := abs(n) - 1
		if nums[i] < 0 {
			o = append(o, i+1)
		}
		nums[i] = -nums[i]
	}
	return o
}
