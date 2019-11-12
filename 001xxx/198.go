import "sort"

func smallestCommonElement(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	for i := 0; i < n; i++ {
		b, num := true, mat[0][i]
		for j := 1; j < m; j++ {
			r := mat[j]
			h := sort.Search(len(r), func(k int) bool {
				return r[k] >= num
			})
			if h == len(r) || r[h] != num {
				b = false
				break
			}
		}
		if b {
			return num
		}
	}
	return -1
}
