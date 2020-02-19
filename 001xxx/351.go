import "sort"

func countNegatives(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	o, k := 0, n
	for i := 0; i < m; i++ {
		ar := grid[i]
		k = sort.Search(k, func(j int) bool { return ar[j] < 0 })
		if k == 0 {
			o += n * (m - i)
			break
		} else {
			o += n - k
		}
	}
	return o
}
