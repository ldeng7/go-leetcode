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

func numSimilarGroups(A []string) int {
	m := map[string]bool{}
	for _, s := range A {
		m[s] = true
	}
	i, ss := 0, make([]string, len(m))
	for s, _ := range m {
		ss[i], i = s, i+1
	}

	l := len(ss)
	if l <= 1 {
		return 1
	}
	au := (&ArrayUnion{}).Init(l)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			r1, r2 := au.GetRoot(i), au.GetRoot(j)
			if r1 == r2 {
				continue
			}
			sa, sb, c := ss[i], ss[j], 0
			for k := 0; k < len(sa); k++ {
				if sa[k] != sb[k] {
					c++
					if c > 2 {
						break
					}
				}
			}
			if c <= 2 {
				au.Set(r1, r2)
			}
		}
	}
	o := 0
	for i := 0; i < l; i++ {
		if au.Get(i) == -1 {
			o++
		}
	}
	return o
}
