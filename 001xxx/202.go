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

func (au *ArrayUnion) GetRoot(i int) int {
	for {
		r := au.arr[i]
		if r == i || r == -1 {
			return i
		}
		i = r
	}
}

func (au *ArrayUnion) Union(a, b int) bool {
	ra, rb := au.GetRoot(a), au.GetRoot(b)
	if ra == rb {
		return false
	}
	au.arr[ra] = rb
	return true
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	o, l := []byte(s), len(s)
	au := (&ArrayUnion{}).Init(l)
	for _, p := range pairs {
		au.Union(p[0], p[1])
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
