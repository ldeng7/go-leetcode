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
	for {
		r := au.arr[i]
		if r == i || r == -1 {
			return i
		}
		i = r
	}
}

func (au *ArrayUnion) Union(a, b int) bool {
	ra, rb := au.GetRoot(a), au.GetRoot(b)
	if ra == rb {
		return false
	}
	au.arr[ra] = rb
	return true
}

func (au *ArrayUnion) CountRoot() int {
	o := 0
	for i, r := range au.arr {
		if r == i || r == -1 {
			o++
		}
	}
	return o
}

func makeConnected(n int, connections [][]int) int {
	if n > len(connections)+1 {
		return -1
	}
	au := (&ArrayUnion{}).Init(n)
	for _, c := range connections {
		au.Union(c[0], c[1])
	}
	return au.CountRoot() - 1
}
