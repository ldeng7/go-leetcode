import "sort"

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

func (au *ArrayUnion) SetRoot(i, v int) {
	au.arr[i] = au.GetRoot(v)
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	o, l := []byte(s), len(s)
	au := (&ArrayUnion{}).Init(l)
	for _, p := range pairs {
		r0, r1 := au.GetRoot(p[0]), au.GetRoot(p[1])
		if r0 != r1 {
			au.Set(r0, r1)
		}
	}
	g := map[int][]int{}
	for i := 0; i < l; i++ {
		r := au.GetRoot(i)
		g[r] = append(g[r], i)
	}
	for _, ar := range g {
		bs := make([]byte, 0, len(ar))
		for _, i := range ar {
			bs = append(bs, s[i])
		}
		sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
		for i, c := range bs {
			o[ar[i]] = c
		}
	}
	return string(o)
}
