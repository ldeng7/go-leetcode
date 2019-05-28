func getModifiedArray(length int, updates [][]int) []int {
	out := make([]int, length+1)
	for _, u := range updates {
		out[u[0]] += u[2]
		out[u[1]+1] -= u[2]
	}
	for i := 1; i <= length; i++ {
		out[i] += out[i-1]
	}
	return out[:length]
}
