func canChoose(groups [][]int, nums []int) bool {
	m, n := len(groups), len(nums)
	i := 0
	for j := 0; i < m && j < n; {
		i1, j1 := 0, j
		g := groups[i]
		l := len(g)
		for ; i1 < l && j1 < n; i1, j1 = i1+1, j1+1 {
			if g[i1] != nums[j1] {
				break
			}
		}
		if i1 == l {
			i, j = i+1, j1
		} else {
			j++
		}
	}
	return i == m
}
