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

func avoidFlood(rains []int) []int {
	l := len(rains)
	o, ar := make([]int, l), make([]int, l)
	m := map[int]int{}
	au := (&ArrayUnion{}).Init(l)

	for i, r := range rains {
		if r == 0 {
			o[i] = 1
			if i != 0 && rains[i-1] > 0 {
				ar[au.GetRoot(i-1)] = i
			}
			continue
		}

		if i != 0 && rains[i-1] > 0 {
			au.Set(i, au.Get(i-1))
		} else {
			au.Set(i, i)
			ar[i] = 100001
		}

		if v, ok := m[r]; ok {
			rt := au.GetRoot(v)
			j := ar[rt]
			if j >= i {
				return []int{}
			}
			o[j] = r
			j++
			ar[rt] = j
			if j < l && rains[j] > 0 {
				au.Set(rt, au.GetRoot(j))
			}
		}
		o[i], m[r] = -1, i
	}
	return o
}
