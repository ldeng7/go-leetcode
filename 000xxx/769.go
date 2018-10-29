func maxChunksToSorted(arr []int) int {
	out := 0
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > j {
			j = arr[i]
		}
		if i == j {
			out++
		}
	}
	return out
}
