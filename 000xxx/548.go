func splitArray(nums []int) bool {
	if len(nums) < 7 {
		return false
	}
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	for i := 3; i < n-3; i++ {
		m := map[int]bool{}
		for j := 1; j < i-1; j++ {
			if nums[j-1] == nums[i-1]-nums[j] {
				m[nums[j-1]] = true
			}
		}
		for j := i + 1; j < n-1; j++ {
			s3, s4 := nums[j-1]-nums[i], nums[n-1]-nums[j]
			if s3 == s4 && m[s3] {
				return true
			}
		}
	}
	return false
}
