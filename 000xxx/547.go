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

func findCircleNum(M [][]int) int {
	l := len(M)
	o := l
	au := (&ArrayUnion{}).Init(l)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if M[i][j] == 1 && au.Union(j, i) {
				o--
			}
		}
	}
	return o
}
