func findLengthOfLCIS(nums []int) int {
	if 0 == len(nums) {
		return 0
	} else if 1 == len(nums) {
		return 1
	}
	i, j := 1, 0
	out := 1
	for ; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			j = i
			continue
		}
		l := i - j + 1
		if l > out {
			out = l
		}
	}
	return out
}
