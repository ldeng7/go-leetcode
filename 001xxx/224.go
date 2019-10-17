func maxEqualFreq(nums []int) int {
	c, f := [100001]int{}, [100001]int{}
	o, m := 0, 0
	for i, n := range nums {
		c[n]++
		cn := c[n]
		f[cn]++
		if cn > m {
			m = cn
		}
		if (f[m] == 1 && (f[m-1]-1)*(m-1)+m == i+1) || f[m]*m == i {
			o = i + 1
		}
	}
	if 1 == m {
		return len(nums)
	}
	return o
}
