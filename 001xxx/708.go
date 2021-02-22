func largestSubarray(nums []int, k int) []int {
	i, l, m := 0, len(nums), nums[0]
	for j := 1; j <= l-k; j++ {
		if v := nums[j]; v > m {
			m, i = v, j
		}
	}
	return nums[i : i+k]
}
