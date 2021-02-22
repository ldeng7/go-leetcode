func maxResult(nums []int, k int) int {
	l := len(nums)
	q := make([]int, 1, 16)
	q[0] = 0
	for i := 1; i < l; i++ {
		for 0 != len(q) && q[0]+k < i {
			q = q[1:]
		}
		nums[i] += nums[q[0]]
		for 0 != len(q) && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	return nums[l-1]
}
