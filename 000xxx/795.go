func numSubarrayBoundedMax(A []int, L int, R int) int {
	o, l, r := 0, -1, -1
	for i, a := range A {
		if a > R {
			l, r = i, i
			continue
		}
		if a >= L {
			r = i
		}
		o += r - l
	}
	return o
}
