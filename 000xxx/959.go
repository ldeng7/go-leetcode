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

func (au *ArrayUnion) GetRoot(i int) int {
	for {
		r := au.arr[i]
		if r == i || r == -1 {
			return i
		}
		i = r
	}
}

func (au *ArrayUnion) SetRoot(i, v int) {
	au.arr[i] = au.GetRoot(v)
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
				au.SetRoot(au.GetRoot(b), b+3)
				au.SetRoot(au.GetRoot(b+1), b+2)
			case '\\':
				au.SetRoot(au.GetRoot(b), b+1)
				au.SetRoot(au.GetRoot(b+2), b+3)
			case ' ':
				au.SetRoot(au.GetRoot(b), b+1)
				au.SetRoot(au.GetRoot(b+1), b+2)
				au.SetRoot(au.GetRoot(b+2), b+3)
			}
			if i > 0 {
				au.SetRoot(au.GetRoot(b), b-(l<<2)+2)
			}
			if j > 0 {
				au.SetRoot(au.GetRoot(b+3), b-3)
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
