func findMin(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] <= nums[j] {
			return nums[i]
		}
		if j == i+1 {
			return nums[j]
		}
		h := i + (j-i)>>1
		if nums[h] < nums[j] {
			j = h
		} else {
			i = h
		}
	}
	return -1
}
