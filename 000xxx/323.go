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

func countComponents(n int, edges [][]int) int {
	o := n
	au := (&ArrayUnion{}).Init(n)
	for _, e := range edges {
		if au.Union(e[1], e[0]) {
			o--
		}
	}
	return o
}
