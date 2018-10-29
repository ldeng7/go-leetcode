func jump(nums []int) int {
	out, j, k := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > j {
			j = i + nums[i]
		}
		if i == k {
			k = j
			out++
			if j >= len(nums)-1 {
				break
			}
		}
	}
	return out
}
