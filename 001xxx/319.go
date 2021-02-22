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

func (au *ArrayUnion) Get(i int) int {
	return au.arr[i]
}

func (au *ArrayUnion) Merge(i, j int) bool {
	r1, r2 := au.GetRoot(i), au.GetRoot(j)
	if r1 != r2 {
		au.arr[r1] = r2
		return true
	}
	return false
}

func makeConnected(n int, connections [][]int) int {
	if n > len(connections)+1 {
		return -1
	}
	au := (&ArrayUnion{}).Init(n)
	for _, c := range connections {
		au.Merge(c[0], c[1])
	}
	o := 0
	for i := 0; i < n; i++ {
		if r := au.Get(i); r == -1 || r == i {
			o++
		}
	}
	return o - 1
}
