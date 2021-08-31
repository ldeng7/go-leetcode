func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxSumMinProduct(nums []int) int {
	l := len(nums)
	ps := make([]int, l+2)
	s := make([]int, 0, 32)
	o := 0
	for i := 0; i <= l; i++ {
		a := 0
		if i < l {
			a = nums[i]
		}
		ps[i+1] = a + ps[i]
		for 0 != len(s) {
			b := nums[s[len(s)-1]]
			if b < a {
				break
			}
			s = s[:len(s)-1]
			j := 0
			if 0 != len(s) {
				j = s[len(s)-1] + 1
			}
			o = max(o, b*(ps[i]-ps[j]))
		}
		s = append(s, i)
	}
	return o % 1000000007
}
