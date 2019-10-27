func circularPermutation(n int, start int) []int {
	o, k := make([]int, 1<<uint(n)), 1
	o[0] = start
	for i := 0; i < n; i++ {
		m := 1 << uint(i)
		for j := k - 1; j >= 0; j-- {
			o[k] = o[j] ^ m
			k++
		}
	}
	return o
}
