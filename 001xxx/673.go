func mostCompetitive(nums []int, k int) []int {
	o, l := make([]int, k), len(nums)
	for i, j := 0, 0; i < l; i++ {
		for ; j != 0 && nums[i] < o[j-1] && k-j < l-i; j-- {
		}
		if j < k {
			o[j] = nums[i]
			j++
		}
	}
	return o
}
