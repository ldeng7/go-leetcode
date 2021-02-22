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

func (au *ArrayUnion) GetRoot(i int) int {
	r := au.arr[i]
	if r == -1 || r == i {
		return i
	}
	r = au.GetRoot(r)
	au.arr[i] = r
	return r
}

func (au *ArrayUnion) Merge(i, j int) bool {
	r1, r2 := au.GetRoot(i), au.GetRoot(j)
	if r1 != r2 {
		au.arr[r1] = r2
		return true
	}
	return false
}

func areConnected(n int, threshold int, queries [][]int) []bool {
	o := make([]bool, len(queries))
	if threshold == 0 {
		for i, _ := range queries {
			o[i] = true
		}
		return o
	}
	au := (&ArrayUnion{}).Init(n + 1)
	for i := threshold + 1; i <= n; i++ {
		for j := 1; i*j <= n; j++ {
			au.Merge(i, i*j)
		}
	}
	for i, q := range queries {
		o[i] = au.GetRoot(q[0]) == au.GetRoot(q[1])
	}
	return o
}
