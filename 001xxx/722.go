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

func (au *ArrayUnion) Merge(i, j int) bool {
	r1, r2 := au.GetRoot(i), au.GetRoot(j)
	if r1 != r2 {
		au.arr[r1] = r2
		return true
	}
	return false
}

func (au *ArrayUnion) Get(i int) int {
	return au.arr[i]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	l := len(source)
	au := (&ArrayUnion{}).Init(l)
	for _, p := range allowedSwaps {
		au.Merge(p[0], p[1])
		r := au.GetRoot(p[1])
		au.Set(r, r)
	}

	m := map[int]map[int]int{}
	o := 0
	for i := 0; i < l; i++ {
		if au.Get(i) == -1 {
			if source[i] != target[i] {
				o++
			}
		} else {
			r := au.GetRoot(i)
			mm := m[r]
			if nil == mm {
				mm = map[int]int{}
				m[r] = mm
			}
			mm[source[i]]++
			mm[target[i]]--
		}
	}

	for _, mm := range m {
		for _, c := range mm {
			o += max(0, c)
		}
	}
	return o
}
