func maxChunksToSorted(arr []int) int {
	out := 0
	j := 0
	for i, a := range arr {
		if a > j {
			j = a
		}
		if i == j {
			out++
		}
	}
	return out
}
