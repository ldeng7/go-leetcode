func jump(nums []int) int {
	o, j, k, e := 0, 0, 0, len(nums)-1
	for i := 0; i < e; i++ {
		if i+nums[i] > j {
			j = i + nums[i]
		}
		if i == k {
			k, o = j, o+1
			if j >= e {
				break
			}
		}
	}
	return o
}
