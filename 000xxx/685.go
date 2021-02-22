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

func findRedundantDirectedConnection(edges [][]int) []int {
	l := len(edges)
	au := (&ArrayUnion{}).Init(l + 1)
	var first, second []int
	for _, edge := range edges {
		if au.Get(edge[1]) == -1 {
			au.Set(edge[1], edge[0])
		} else {
			first = []int{au.Get(edge[1]), edge[1]}
			second = []int{edge[0], edge[1]}
			edge[1] = 0
		}
	}
	for i := 0; i <= l; i++ {
		au.Set(i, i)
	}
	for _, edge := range edges {
		if edge[1] == 0 {
			continue
		}
		if !au.Merge(edge[0], edge[1]) {
			if 0 == len(first) {
				return edge
			}
			return first
		}
	}
	return second
}
