func minCostToMoveChips(chips []int) int {
	c0, c1 := 0, 0
	for _, i := range chips {
		if i&1 == 0 {
			c0++
		} else {
			c1++
		}
	}
	if c0 < c1 {
		return c0
	}
	return c1
}
