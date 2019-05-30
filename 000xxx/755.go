func pourWater(heights []int, V int, K int) []int {
	for i := 0; i < V; i++ {
		l, r := K, K
		for l > 0 && heights[l] >= heights[l-1] {
			l--
		}
		for l < K && heights[l] == heights[l+1] {
			l++
		}
		for r < len(heights)-1 && heights[r] >= heights[r+1] {
			r++
		}
		for r > K && heights[r] == heights[r-1] {
			r--
		}
		if heights[l] < heights[K] {
			heights[l]++
		} else {
			heights[r]++
		}
	}
	return heights
}
