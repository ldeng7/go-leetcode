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

func numIslands2(m int, n int, positions [][]int) []int {
	out := []int{}
	c := 0
	au := (&ArrayUnion{}).Init(m * n)

	cal := func(py, px, dy, dx, i int) {
		y, x := py+dy, px+dx
		j := n*y + x
		if y < 0 || y >= m || x < 0 || x >= n || au.Get(j) == -1 {
			return
		}
		r, r1 := au.GetRoot(i), au.GetRoot(j)
		if r != r1 {
			au.Set(r1, r)
			c--
		}
	}

	for _, p := range positions {
		py, px := p[0], p[1]
		i := n*py + px
		if -1 == au.Get(i) {
			au.Set(i, i)
			c++
		}
		cal(py, px, -1, 0, i)
		cal(py, px, 1, 0, i)
		cal(py, px, 0, -1, i)
		cal(py, px, 0, 1, i)
		out = append(out, c)
	}
	return out
}
