func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func sumOfFlooredPairs(nums []int) int {
	m := 0
	for _, a := range nums {
		m = max(m, a)
	}
	o, l := 0, len(nums)
	t := make([]int, m+2)
	for i := 0; i < l; i++ {
		t[nums[i]]++
	}
	for i := m; i > 0; i-- {
		t[i] += t[i+1]
	}
	for i := 1; i <= m; i++ {
		if t[i] <= t[i+1] {
			continue
		}
		s := 0
		for j := i; j <= m; j += i {
			s += t[j]
		}
		o = (o + s*(t[i]-t[i+1])) % 1000000007
	}
	return o
}
