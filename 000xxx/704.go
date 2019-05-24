func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i < j {
		m := int((uint(i) + uint(j)) >> 1)
		if nums[m] >= target {
			j = m
		} else {
			i = m + 1
		}
	}
	if nums[i] != target {
		return -1
	}
	return i
}
