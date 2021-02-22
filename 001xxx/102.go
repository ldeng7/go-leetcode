import "sort"

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

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maximumMinimumPath(A [][]int) int {
	m, n := len(A), len(A[0])
	l := m * n
	au := (&ArrayUnion{}).Init(l)
	sz := make([]int, l)
	for i := 0; i < l; i++ {
		sz[i] = 1
	}
	ar, k := make([][3]int, l), 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ar[k], k = [3]int{A[i][j], i, j}, k+1
		}
	}

	sort.Slice(ar, func(i, j int) bool { return ar[i][0] > ar[j][0] })
	o := min(A[0][0], A[m-1][n-1])
	A[0][0], A[m-1][n-1] = -1, -1
	for au.GetRoot(0) != au.GetRoot(l-1) {
		p, y0, x0 := ar[0][0], ar[0][1], ar[0][2]
		ar = ar[1:]
		o = min(o, p)
		for _, d := range dirs {
			y, x := y0+d[0], x0+d[1]
			if y < 0 || y >= m || x < 0 || x >= n || A[y][x] != -1 {
				continue
			}
			i, j := y0*n+x0, y*n+x
			r1, r2 := au.GetRoot(i), au.GetRoot(j)
			if r1 != r2 {
				if sz[r1] < sz[r2] {
					r1, r2 = r2, r1
				}
				au.Set(r2, r1)
				sz[r1] += sz[r2]
			}
		}
		A[y0][x0] = -1
	}
	return o
}
