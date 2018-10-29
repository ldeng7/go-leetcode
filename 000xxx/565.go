func arrayNesting(nums []int) int {
	out := 0
	for i := 0; i < len(nums); i++ {
		cnt := 1
		for nums[i] != i && nums[i] != nums[nums[i]] {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			cnt++
		}
		if cnt > out {
			out = cnt
		}
	}
	return out
}
