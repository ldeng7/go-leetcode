func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func constrainedSubsetSum(nums []int, k int) int {
	o, l := nums[0], len(nums)
	t := make([]int, l)
	t[0] = o
	q := make([]int, 1, l)
	q[0] = o
	for i := 1; i < l; i++ {
		if i > k && q[0] == t[i-k-1] {
			q = q[1:]
		}
		n := nums[i]
		j := max(n, n+q[0])
		t[i] = j
		for 0 != len(q) && j > q[len(q)-1] {
			q = q[:len(q)-1]
		}
		q = append(q, j)
		o = max(o, j)
	}
	return o
}
