func searchRange(nums []int, target int) []int {
	if 0 == len(nums) {
		return []int{-1, -1}
	}

	i, j := 0, len(nums)
	for i < j {
		h := int((uint(i) + uint(j)) >> 1)
		if nums[h] >= target {
			j = h
		} else {
			i = h + 1
		}
	}
	if i == len(nums) || nums[i] != target {
		return []int{-1, -1}
	}
	out := []int{i, 0}

	j = len(nums)
	for i < j {
		h := int((uint(i) + uint(j)) >> 1)
		if nums[h] > target {
			j = h
		} else {
			i = h + 1
		}
	}
	out[1] = i - 1
	return out
}
