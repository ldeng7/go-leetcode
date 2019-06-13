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

func findCircleNum(M [][]int) int {
	l := len(M)
	o := l
	au := (&ArrayUnion{}).Init(l)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if M[i][j] == 1 {
				r1, r2 := au.GetRoot(i), au.GetRoot(j)
				if r1 != r2 {
					o--
					au.Set(r2, r1)
				}
			}
		}
	}
	return o
}
