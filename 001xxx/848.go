func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func getMinDistance(nums []int, target int, start int) int {
	o := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			o = min(o, abs(i-start))
		}
	}
	return o
}
