import "sort"

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	ar := make([]int, 0, m*n)

	for j, c, r := 0, 0, matrix[0]; j < n; j++ {
		v := r[j]
		c ^= v
		r[j] = c
		ar = append(ar, c)
	}

	for i := 1; i < m; i++ {
		for j, c, r, rp := 0, 0, matrix[i], matrix[i-1]; j < n; j++ {
			v := r[j]
			c ^= v
			v = c ^ rp[j]
			r[j] = v
			ar = append(ar, v)
		}
	}
	sort.Slice(ar, func(i, j int) bool { return ar[i] > ar[j] })
	return ar[k-1]
}
