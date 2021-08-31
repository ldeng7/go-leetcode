func getXORSum(arr1 []int, arr2 []int) int {
	a, b := 0, 0
	for _, n := range arr1 {
		a ^= n
	}
	for _, n := range arr2 {
		b ^= n
	}
	return a & b
}
