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

func regionsBySlashes(grid []string) int {
	l := len(grid)
	lu := l * l << 2
	au := (&ArrayUnion{}).Init(lu)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			b := (i*l + j) << 2
			switch grid[i][j] {
			case '/':
				au.Set(au.GetRoot(b), au.GetRoot(b+3))
				au.Set(au.GetRoot(b+1), au.GetRoot(b+2))
			case '\\':
				au.Set(au.GetRoot(b), au.GetRoot(b+1))
				au.Set(au.GetRoot(b+2), au.GetRoot(b+3))
			case ' ':
				r1, r2 := au.GetRoot(b+1), au.GetRoot(b+2)
				au.Set(au.GetRoot(b), r1)
				au.Set(r1, r2)
				au.Set(r2, au.GetRoot(b+3))
			}
			if i > 0 {
				au.Set(au.GetRoot(b), au.GetRoot(b-(l<<2)+2))
			}
			if j > 0 {
				au.Set(au.GetRoot(b+3), au.GetRoot(b-3))
			}
		}
	}
	o := 0
	for i := 0; i < lu; i++ {
		if r := au.Get(i); r == i || r == -1 {
			o++
		}
	}
	return o
}
