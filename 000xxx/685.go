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

func (au *ArrayUnion) Get(i int) int {
	return au.arr[i]
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
		r1, r2 := au.GetRoot(edge[0]), au.GetRoot(edge[1])
		if r1 == r2 {
			if 0 == len(first) {
				return edge
			}
			return first
		}
		au.Set(r1, r2)
	}
	return second
}
