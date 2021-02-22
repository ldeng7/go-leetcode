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

func findCircleNum(M [][]int) int {
	l := len(M)
	o := l
	au := (&ArrayUnion{}).Init(l)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if M[i][j] == 1 && au.Merge(j, i) {
				o--
			}
		}
	}
	return o
}
