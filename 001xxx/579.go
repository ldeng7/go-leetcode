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

func (au *ArrayUnion) Copy() *ArrayUnion {
	arr := make([]int, len(au.arr))
	copy(arr, au.arr)
	return &ArrayUnion{arr}
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	au := (&ArrayUnion{}).Init(n + 1)
	o, c1 := 0, n
	for _, e := range edges {
		if e[0] == 3 {
			if !au.Merge(e[1], e[2]) {
				o++
			} else {
				c1--
			}
		}
	}
	au1 := au.Copy()
	c2 := c1
	for _, e := range edges {
		if e[0] == 1 {
			if !au.Merge(e[1], e[2]) {
				o++
			} else {
				c1--
			}
		} else if e[0] == 2 {
			if !au1.Merge(e[1], e[2]) {
				o++
			} else {
				c2--
			}
		}
	}
	if c1 != 1 || c2 != 1 {
		return -1
	}
	return o
}
