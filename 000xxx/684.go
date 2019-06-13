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
	for {
		r := au.arr[i]
		if r == i || r == -1 {
			return i
		}
		i = r
	}
}

func findRedundantConnection(edges [][]int) []int {
	au := (&ArrayUnion{}).Init(2001)
	for _, e := range edges {
		r1, r2 := au.GetRoot(e[0]), au.GetRoot(e[1])
		if r1 == r2 {
			return e
		}
		au.Set(r1, r2)
	}
	return []int{}
}
