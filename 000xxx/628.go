func maximumProduct(nums []int) int {
	p1, p2, p3 := -1001, -1001, -1001
	n1, n2 := 1001, 1001
	for _, num := range nums {
		if num > p3 {
			p3 = num
		}
		if p3 > p2 {
			p2, p3 = p3, p2
		}
		if p2 > p1 {
			p1, p2 = p2, p1
		}
		if num < n2 {
			n2 = num
		}
		if n2 < n1 {
			n1, n2 = n2, n1
		}
	}
	a, b := p2*p3, n1*n2
	if a > b {
		return p1 * a
	} else {
		return p1 * b
	}
}
