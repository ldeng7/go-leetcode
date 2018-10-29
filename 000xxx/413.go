func numberOfArithmeticSlices(a []int) int {
	out, cur := 0, 0
	for i := 2; i < len(a); i++ {
		if a[i]-a[i-1] == a[i-1]-a[i-2] {
			cur++
			out += cur
		} else {
			cur = 0
		}
	}
	return out
}
