func countTriplets(arr []int) int {
	o, l := 0, len(arr)
	t := make([]int, l+1)
	for i, n := range arr {
		t[i+1] = n ^ t[i]
	}
	for i := 0; i < l; i++ {
		xor := t[i]
		for j := i + 1; j < l; j++ {
			if xor == t[j+1] {
				o += j - i
			}
		}
	}
	return o
}
