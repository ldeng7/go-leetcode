func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func maxValueAfterReverse(nums []int) int {
	o, l := 0, len(nums)
	if l <= 1 {
		return 0
	}
	p := nums[0]
	for i := 1; i < l; i++ {
		p1 := nums[i]
		o += abs(p - p1)
		p = p1
	}
	s, n0, nl := o, nums[0], nums[l-1]
	for i := 1; i <= l-2; i++ {
		np, n, nn := nums[i-1], nums[i], nums[i+1]
		o = max(o, max(s+abs(nl-np)-abs(n-np), s+abs(nn-n0)-abs(nn-n)))
	}
	var m1, m2, m3, m4 int = -1e9, -1e9, -1e9, -1e9
	for i := 1; i < l; i++ {
		n, np := nums[i], nums[i-1]
		sum, dif := n+np, n-np
		ab := abs(dif)
		o = max(o, max(max(s+m1+sum-ab, s+m2-sum-ab), max(s+m3+dif-ab, s+m4-dif-ab)))
		m1, m2, m3, m4 = max(m1, -sum-ab), max(m2, sum-ab), max(m3, -dif-ab), max(m4, dif-ab)
	}
	return o
}
