type ArrayUnion struct {
	arr []int
}

func (au *ArrayUnion) Init(l int) *ArrayUnion {
	au.arr = make([]int, l)
	for i := 0; i < l; i++ {
		au.arr[i] = -1
	}
	return au
}

func (au *ArrayUnion) Set(i, v int) {
	au.arr[i] = v
}

func (au *ArrayUnion) GetRoot(i int) int {
	r := au.arr[i]
	if r == -1 || r == i {
		return i
	}
	r = au.GetRoot(r)
	au.arr[i] = r
	return r
}

func containsCycle(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	au := (&ArrayUnion{}).Init(m * n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			k := i*n + j
			if i > 0 && grid[i-1][j] == grid[i][j] {
				au.Set(k, au.GetRoot(k-n))
			}
			if j > 0 && grid[i][j-1] == grid[i][j] {
				r1, r2 := au.GetRoot(k-1), au.GetRoot(k)
				if r1 == r2 {
					return true
				}
				au.Set(r1, r2)
			}
		}
	}
	return false
}
