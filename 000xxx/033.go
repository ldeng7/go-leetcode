func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		h := i + (j-i)>>1
		if nums[h] == target {
			return h
		}
		if nums[h] < nums[j] {
			if nums[h] < target && target <= nums[j] {
				i = h + 1
			} else {
				j = h - 1
			}
		} else {
			if nums[i] <= target && target < nums[h] {
				j = h - 1
			} else {
				i = h + 1
			}
		}
	}
	return -1
}
