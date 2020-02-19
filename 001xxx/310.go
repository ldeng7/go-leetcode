func xorQueries(arr []int, queries [][]int) []int {
	v := make([]int, len(arr)+1)
	for i, n := range arr {
		v[i+1] = v[i] ^ n
	}
	o := make([]int, len(queries))
	for i, q := range queries {
		o[i] = v[q[1]+1] ^ v[q[0]]
	}
	return o
}
