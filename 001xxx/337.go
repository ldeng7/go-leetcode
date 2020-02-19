import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	m, n := len(mat), len(mat[0])
	o, t := make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		o[i] = i
		ar := mat[i]
		t[i] = sort.Search(n, func(j int) bool { return ar[j] == 0 })
	}
	sort.Slice(o, func(i, j int) bool {
		a, b := o[i], o[j]
		return t[a] < t[b] || (t[a] == t[b] && a < b)
	})
	return o[:k]
}
