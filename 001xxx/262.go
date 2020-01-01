func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxSumDivThree(nums []int) int {
	t := [3]int{0, 0, 0}
	for _, n := range nums {
		m := n % 3
		a, b, c := t[(3-m)%3], t[(4-m)%3], t[(5-m)%3]
		if a != 0 || m == 0 {
			t[0] = max(t[0], a+n)
		}
		if b != 0 || m == 1 {
			t[1] = max(t[1], b+n)
		}
		if c != 0 || m == 2 {
			t[2] = max(t[2], c+n)
		}
	}
	return t[0]
}
